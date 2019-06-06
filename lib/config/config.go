package config

import (
	"github.com/spf13/viper"
	"log"
)

var Email163Config *EmailConfig

func init() {
	if Email163Config == nil {
		viper.AddConfigPath("../config")
		viper.SetConfigName("163email")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalln("加载不到163email的配置:", err)
		}

		Email163Config = new(EmailConfig)
		Email163Config.Host = viper.GetString("Host")
		Email163Config.Port = viper.GetInt("Port")
		Email163Config.Email = viper.GetString("Email")
		Email163Config.Passwd = viper.GetString("Passwd")
	}

}

//发送email的配置
type EmailConfig struct {
	//服务器地址
	Host string
	//端口
	Port int
	//发送人的email
	Email string
	//发送人的密码
	Passwd string
}
