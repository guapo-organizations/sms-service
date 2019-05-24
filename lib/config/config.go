package config

var Email163Config *EmailConfig

func init() {
	if Email163Config == nil {
		Email163Config = new(EmailConfig)
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
