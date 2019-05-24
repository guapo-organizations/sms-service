package sms

import (
	"fmt"
	"github.com/go-mail/mail"
	"github.com/go-redis/redis"
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/help"
	"github.com/guapo-organizations/sms-service/lib/config"
	"github.com/guapo-organizations/sms-service/lib/limit"
	"github.com/guapo-organizations/sms-service/lib/msg"
	"github.com/guapo-organizations/sms-service/lib/template"
	"github.com/lifei6671/gorand"
	"log"
	"strings"
	"time"
)

//邮箱发送验证码
func SendEmailCode(email string) (bool, error) {

	if !help.VerifyEmailFormat(email) {
		return false, fmt.Errorf("邮箱格式不对")
	}

	redis_client := cache.GetRedisClient()
	//如果err没有报错，则是获取到了值
	_, err := redis_client.Get(limit.CreateLimiKey(email, msg.PUBLISH_MODE_EMAIL, msg.PUBLISH_TYPE_CODE)).Result()
	if err == redis.Nil {
		//邮箱没有发送过
		//生成验证码
		code := gorand.KRand(6, gorand.KC_RAND_KIND_ALL)
		//发送验证码
		d := mail.NewDialer(config.Email163Config.Host, config.Email163Config.Port, config.Email163Config.Email, config.Email163Config.Passwd)
		m := mail.NewMessage()
		m.SetHeader("From", config.Email163Config.Email)
		m.SetHeader("To", email)
		m.SetHeader("Subject", template.Code163Template.Subject)
		m.SetBody("text/html", fmt.Sprintf(template.Code163Template.Body, code))
		err = d.DialAndSend(m)
		log.Println("邮箱错误信息是什么:", err)

		if err != nil {
			return false, err
		}

		//设置验证码缓存,为1分钟
		_, err = redis_client.SetNX(limit.CreateLimiKey(email, msg.PUBLISH_MODE_EMAIL, msg.PUBLISH_TYPE_CODE), code, 60*time.Second).Result()

		if err != nil {
			return false, err
		}

		//发送成功
		return true, nil
	}

	if err != nil {
		//也可能是连接出错了
		return false, err
	}

	//邮箱已经发送
	return false, fmt.Errorf("我丢，你已经发送过验证码了，稍后再发好吧？！")

}

//验证code是否正确
func ValidateCode(key string, publish_mode string, publish_type string) (bool, error) {
	redis_client := cache.GetRedisClient()
	result, err := redis_client.Get(limit.CreateLimiKey(key, publish_mode, publish_type)).Result()
	if err == redis.Nil {
		return false, fmt.Errorf("%s 验证码过期", key)
	}

	//可能连接出错
	if err != nil {
		return false, err
	}

	//比较一下
	if !strings.EqualFold(result, key) {
		return false, nil
	}
	
	return true, nil
}
