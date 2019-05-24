package template

var Code163Template *CodeTemplate

func init() {
	if Code163Template == nil {
		Code163Template = new(CodeTemplate)
		Code163Template.Subject = "最靓的仔"
		Code163Template.Body = "亲，您有一份万事屋工作职位邀请码哦~：%s"
	}
}


type CodeTemplate struct {
	Subject string
	Body    string
}
