package event_message

// 时间model对象
type EventMessage struct {
	MessageId        string `json:"messageId"`        // 消息队列的消息id
	EventMessageType string `json:"eventMessageType"` // 事件类型
	BizId            string `json:"bizId"`            // 业务id
	AccountNo        int64  `json:"accountNo"`        // 账号
	Content          string `json:"content"`          // 消息体
	Remark           string `json:"remark"`           // 备注
}
