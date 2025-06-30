package event

// 处理发送短信事件
func (e EventHandler) HandleSendSmsEvent(data interface{}) {
	v := data.(*events.SendSmsEvent)
	if v != nil {
		v.Template = sms.ResolveMessage(v.Template, v.Data)
		ev := &proto.EVSendSmsEventData{
			Provider:     int32(v.Provider),
			Phone:        v.Phone,
			Template:     v.Template,
			TemplateCode: v.TemplateCode,
			SpTemplateId: v.SpTemplateId,
			Data:         v.Data,
		}
		msq.Push(msq.SendSmsTopic, typeconv.MustJson(ev))
	}
}
