package template

import (
	"github.com/spf13/viper"
	"log"
)

var Code163Template *CodeTemplate

func init() {
	if Code163Template == nil {
		viper.AddConfigPath("../config")
		viper.SetConfigName("163email")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalln("template模板加载配置失败")
		}
		Code163Template = new(CodeTemplate)
		Code163Template.Subject = viper.GetString("Subject")
		Code163Template.Body = viper.GetString("Body")
	}
}


type CodeTemplate struct {
	Subject string
	Body    string
}
