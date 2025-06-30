package event

// 发送短信事件
type SendSmsEvent struct {
	// 短信服务商
	Provider int
	// 手机号
	Phone string
	// 短信内容
	Template string
	// 模板代码
	TemplateCode string
	// 短信模板ID
	SpTemplateId string
	// 数据
	Data []string
}
