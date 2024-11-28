package enums

// TransferStatusNameCN 用于映射枚举值到中文注释
var (
	TransferStatus_TransferStatus_CN = map[TransferStatus_TransferStatus]string{
		TransferStatus_PENDING: "未转账",
		TransferStatus_SUCCESS: "已转账",
		TransferStatus_FAIL:    "转账失败",
	}
	//支付
	OrderState_OrderState_CN = map[OrderState_OrderState]string{
		OrderState_NEW:       "已支付",
		OrderState_COMPLETE:  "完成",
		OrderState_CANCELLED: "取消",
		OrderState_FAIL:      "失败",
	}

	BillingTypeEnums_CN = map[BillingTypeEnums]string{
		BillingTypeEnums_AMOUNT: "按量",
		BillingTypeEnums_DAY:    "按日",
		BillingTypeEnums_MONTH:  "按月",
		BillingTypeEnums_WEEK:   "按周",
		BillingTypeEnums_YEAR:   "按年",
		BillingTypeEnums_RENEW:  "续费",
	}

	//支付状态
	PayState_PayState_CN = map[PayState_PayState]string{
		PayState_PAY:        "已支付",
		PayState_FAILED:     "失败",
		PayState_REFUNDED:   "已退款",
		PayState_CANCELLED:  "已取消",
		PayState_PROCESSING: "处理中",
		PayState_NEW:        "新订单",
		PayState_CANCEL:     "超时取消",
		PayState_FAIL:       "创建失败",
	}

	OrderType_OrderTypeEnums_CN = map[OrderType_OrderTypeEnums]string{
		OrderType_AMOUNT: "按量",
		OrderType_DAY:    "按日",
		OrderType_MONTH:  "按月",
		OrderType_WEEK:   "按周",
		OrderType_YEAR:   "按年",
		OrderType_RENEW:  "续费",
	}
	MessageType_MessageType_CN = map[MessageType_MessageType]string{
		MessageType_TEXT:     "文本",
		MessageType_FILE:     "文件",
		MessageType_IMAGE:    "图片",
		MessageType_INSTANCE: "实例",
	}
	PayType_CN = map[PayStrategy_PayType]string{
		PayStrategy_ALI:      "支付宝",
		PayStrategy_WECHAT:   "微信",
		PayStrategy_INTEGRAL: "雨点",
	}

	SettlementApplication_SettlementStatus_CN = map[SettlementApplication_SettlementStatus]string{
		//待处理
		SettlementApplication_PENDING: "待处理",
		//已结算
		SettlementApplication_SETTLED: "已结算",
		//拒绝
		SettlementApplication_REJECTED: "拒绝",
	}
)
