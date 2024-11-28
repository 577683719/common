package enums_v2

var (
	SettlementApplication_SettlementStatus_name_desc = map[int32]string{
		0: "待处理",
		1: "已结算",
		2: "拒绝",
	}

	SettlementApplication_SettlementStatus_value_desc = map[string]string{
		"PENDING":  "待处理",
		"SETTLED":  "已结算",
		"REJECTED": "拒绝",
	}
)

func NewSettlementApplication_SettlementStatusFromNumber(number int32) *SettlementApplication_SettlementStatus {
	e := SettlementApplication_SettlementStatus(number)
	return &e
}

func NewSettlementApplication_SettlementStatusFromValue(value string) *SettlementApplication_SettlementStatus {
	return NewSettlementApplication_SettlementStatusFromNumber(SettlementApplication_SettlementStatus_value[value])
}

func (e SettlementApplication_SettlementStatus) GetDesc() string {
	return SettlementApplication_SettlementStatus_name_desc[int32(e)]
}

func (SettlementApplication_SettlementStatus) GetDescFromNumber(number int32) string {
	return SettlementApplication_SettlementStatus_name_desc[number]
}

func (SettlementApplication_SettlementStatus) GetDescFromName(name string) string {
	return SettlementApplication_SettlementStatus_value_desc[name]
}

var (
	TransferStatus_TransferStatus_name_desc = map[int32]string{
		0: "未转账",
		1: "已转账",
		2: "转账失败",
	}

	TransferStatus_TransferStatus_value_desc = map[string]string{
		"PENDING": "未转账",
		"SUCCESS": "已转账",
		"FAIL":    "转账失败",
	}
)

func NewTransferStatus_TransferStatusFromNumber(number int32) *TransferStatus_TransferStatus {
	e := TransferStatus_TransferStatus(number)
	return &e
}

func NewTransferStatus_TransferStatusFromValue(value string) *TransferStatus_TransferStatus {
	return NewTransferStatus_TransferStatusFromNumber(TransferStatus_TransferStatus_value[value])
}

func (e TransferStatus_TransferStatus) GetDesc() string {
	return TransferStatus_TransferStatus_name_desc[int32(e)]
}

func (TransferStatus_TransferStatus) GetDescFromNumber(number int32) string {
	return TransferStatus_TransferStatus_name_desc[number]
}

func (TransferStatus_TransferStatus) GetDescFromName(name string) string {
	return TransferStatus_TransferStatus_value_desc[name]
}

var (
	ApprovalStatus_ApprovalStatus_name_desc = map[int32]string{
		0: "待审批",
		1: "已批准",
		2: "已拒绝",
		3: "审核中",
		4: "已完成",
		5: "正在审查",
		6: "已取消",
	}

	ApprovalStatus_ApprovalStatus_value_desc = map[string]string{
		"PENDING":      "待审批",
		"APPROVED":     "已批准",
		"REJECTED":     "已拒绝",
		"IN_REVIEW":    "审核中",
		"COMPLETED":    "已完成",
		"UNDER_REVIEW": "正在审查",
		"CANCELLED":    "已取消",
	}
)

func NewApprovalStatus_ApprovalStatusFromNumber(number int32) *ApprovalStatus_ApprovalStatus {
	e := ApprovalStatus_ApprovalStatus(number)
	return &e
}

func NewApprovalStatus_ApprovalStatusFromValue(value string) *ApprovalStatus_ApprovalStatus {
	return NewApprovalStatus_ApprovalStatusFromNumber(ApprovalStatus_ApprovalStatus_value[value])
}

func (e ApprovalStatus_ApprovalStatus) GetDesc() string {
	return ApprovalStatus_ApprovalStatus_name_desc[int32(e)]
}

func (ApprovalStatus_ApprovalStatus) GetDescFromNumber(number int32) string {
	return ApprovalStatus_ApprovalStatus_name_desc[number]
}

func (ApprovalStatus_ApprovalStatus) GetDescFromName(name string) string {
	return ApprovalStatus_ApprovalStatus_value_desc[name]
}

var (
	RechargeType_RechargeTypeEnum_name_desc = map[int32]string{
		0: "雨点充值",
		1: "余额充值",
	}

	RechargeType_RechargeTypeEnum_value_desc = map[string]string{
		"INTEGRAL": "雨点充值",
		"WALLET":   "余额充值",
	}
)

func NewRechargeType_RechargeTypeEnumFromNumber(number int32) *RechargeType_RechargeTypeEnum {
	e := RechargeType_RechargeTypeEnum(number)
	return &e
}

func NewRechargeType_RechargeTypeEnumFromValue(value string) *RechargeType_RechargeTypeEnum {
	return NewRechargeType_RechargeTypeEnumFromNumber(RechargeType_RechargeTypeEnum_value[value])
}

func (e RechargeType_RechargeTypeEnum) GetDesc() string {
	return RechargeType_RechargeTypeEnum_name_desc[int32(e)]
}

func (RechargeType_RechargeTypeEnum) GetDescFromNumber(number int32) string {
	return RechargeType_RechargeTypeEnum_name_desc[number]
}

func (RechargeType_RechargeTypeEnum) GetDescFromName(name string) string {
	return RechargeType_RechargeTypeEnum_value_desc[name]
}

var (
	MessageType_MessageType_name_desc = map[int32]string{
		0: "普通文本消息",
		1: "图片消息",
		2: "文件消息",
		3: "实例消息",
	}

	MessageType_MessageType_value_desc = map[string]string{
		"TEXT":     "普通文本消息",
		"IMAGE":    "图片消息",
		"FILE":     "文件消息",
		"INSTANCE": "实例消息",
	}
)

func NewMessageType_MessageTypeFromNumber(number int32) *MessageType_MessageType {
	e := MessageType_MessageType(number)
	return &e
}

func NewMessageType_MessageTypeFromValue(value string) *MessageType_MessageType {
	return NewMessageType_MessageTypeFromNumber(MessageType_MessageType_value[value])
}

func (e MessageType_MessageType) GetDesc() string {
	return MessageType_MessageType_name_desc[int32(e)]
}

func (MessageType_MessageType) GetDescFromNumber(number int32) string {
	return MessageType_MessageType_name_desc[number]
}

func (MessageType_MessageType) GetDescFromName(name string) string {
	return MessageType_MessageType_value_desc[name]
}

var (
	TransactionChannel_TransactionChannel_name_desc = map[int32]string{
		0: "雨点",
		1: "支付宝",
		2: "微信",
		3: "钱包",
	}

	TransactionChannel_TransactionChannel_value_desc = map[string]string{
		"INTEGRAL": "雨点",
		"ALI":      "支付宝",
		"WECHAT":   "微信",
		"WALLET":   "钱包",
	}
)

func NewTransactionChannel_TransactionChannelFromNumber(number int32) *TransactionChannel_TransactionChannel {
	e := TransactionChannel_TransactionChannel(number)
	return &e
}

func NewTransactionChannel_TransactionChannelFromValue(value string) *TransactionChannel_TransactionChannel {
	return NewTransactionChannel_TransactionChannelFromNumber(TransactionChannel_TransactionChannel_value[value])
}

func (e TransactionChannel_TransactionChannel) GetDesc() string {
	return TransactionChannel_TransactionChannel_name_desc[int32(e)]
}

func (TransactionChannel_TransactionChannel) GetDescFromNumber(number int32) string {
	return TransactionChannel_TransactionChannel_name_desc[number]
}

func (TransactionChannel_TransactionChannel) GetDescFromName(name string) string {
	return TransactionChannel_TransactionChannel_value_desc[name]
}

var (
	TransactionType_TransactionType_name_desc = map[int32]string{
		0: "消费",
		1: "充值",
	}

	TransactionType_TransactionType_value_desc = map[string]string{
		"CONSUMPTION": "消费",
		"RECHARGE":    "充值",
	}
)

func NewTransactionType_TransactionTypeFromNumber(number int32) *TransactionType_TransactionType {
	e := TransactionType_TransactionType(number)
	return &e
}

func NewTransactionType_TransactionTypeFromValue(value string) *TransactionType_TransactionType {
	return NewTransactionType_TransactionTypeFromNumber(TransactionType_TransactionType_value[value])
}

func (e TransactionType_TransactionType) GetDesc() string {
	return TransactionType_TransactionType_name_desc[int32(e)]
}

func (TransactionType_TransactionType) GetDescFromNumber(number int32) string {
	return TransactionType_TransactionType_name_desc[number]
}

func (TransactionType_TransactionType) GetDescFromName(name string) string {
	return TransactionType_TransactionType_value_desc[name]
}

var (
	IncomeExpenseType_IncomeExpenseType_name_desc = map[int32]string{
		0: "支出",
		1: "收入",
	}

	IncomeExpenseType_IncomeExpenseType_value_desc = map[string]string{
		"EXPENDITURE": "支出",
		"INCOME":      "收入",
	}
)

func NewIncomeExpenseType_IncomeExpenseTypeFromNumber(number int32) *IncomeExpenseType_IncomeExpenseType {
	e := IncomeExpenseType_IncomeExpenseType(number)
	return &e
}

func NewIncomeExpenseType_IncomeExpenseTypeFromValue(value string) *IncomeExpenseType_IncomeExpenseType {
	return NewIncomeExpenseType_IncomeExpenseTypeFromNumber(IncomeExpenseType_IncomeExpenseType_value[value])
}

func (e IncomeExpenseType_IncomeExpenseType) GetDesc() string {
	return IncomeExpenseType_IncomeExpenseType_name_desc[int32(e)]
}

func (IncomeExpenseType_IncomeExpenseType) GetDescFromNumber(number int32) string {
	return IncomeExpenseType_IncomeExpenseType_name_desc[number]
}

func (IncomeExpenseType_IncomeExpenseType) GetDescFromName(name string) string {
	return IncomeExpenseType_IncomeExpenseType_value_desc[name]
}

var (
	OrderState_OrderState_name_desc = map[int32]string{
		0: "已支付",
		1: "完成",
		2: "取消",
		4: "失败",
	}

	OrderState_OrderState_value_desc = map[string]string{
		"NEW":       "已支付",
		"COMPLETE":  "完成",
		"CANCELLED": "取消",
		"FAIL":      "失败",
	}
)

func NewOrderState_OrderStateFromNumber(number int32) *OrderState_OrderState {
	e := OrderState_OrderState(number)
	return &e
}

func NewOrderState_OrderStateFromValue(value string) *OrderState_OrderState {
	return NewOrderState_OrderStateFromNumber(OrderState_OrderState_value[value])
}

func (e OrderState_OrderState) GetDesc() string {
	return OrderState_OrderState_name_desc[int32(e)]
}

func (OrderState_OrderState) GetDescFromNumber(number int32) string {
	return OrderState_OrderState_name_desc[number]
}

func (OrderState_OrderState) GetDescFromName(name string) string {
	return OrderState_OrderState_value_desc[name]
}

var (
	PayState_PayState_name_desc = map[int32]string{
		0: "已支付",
		1: "失败",
		2: "已退款",
		3: "已取消",
		4: "处理中",
		5: "未支付订单",
		6: "超时取消订单",
		7: "创建失败",
	}

	PayState_PayState_value_desc = map[string]string{
		"PAY":        "已支付",
		"FAILED":     "失败",
		"REFUNDED":   "已退款",
		"CANCELLED":  "已取消",
		"PROCESSING": "处理中",
		"NEW":        "未支付订单",
		"CANCEL":     "超时取消订单",
		"FAIL":       "创建失败",
	}
)

func NewPayState_PayStateFromNumber(number int32) *PayState_PayState {
	e := PayState_PayState(number)
	return &e
}

func NewPayState_PayStateFromValue(value string) *PayState_PayState {
	return NewPayState_PayStateFromNumber(PayState_PayState_value[value])
}

func (e PayState_PayState) GetDesc() string {
	return PayState_PayState_name_desc[int32(e)]
}

func (PayState_PayState) GetDescFromNumber(number int32) string {
	return PayState_PayState_name_desc[number]
}

func (PayState_PayState) GetDescFromName(name string) string {
	return PayState_PayState_value_desc[name]
}

var (
	OrderType_OrderTypeEnums_name_desc = map[int32]string{
		0: "按量",
		1: "按日",
		2: "按月",
		3: "周",
		4: "年",
		5: "续费",
	}

	OrderType_OrderTypeEnums_value_desc = map[string]string{
		"AMOUNT": "按量",
		"DAY":    "按日",
		"MONTH":  "按月",
		"WEEK":   "周",
		"YEAR":   "年",
		"RENEW":  "续费",
	}
)

func NewOrderType_OrderTypeEnumsFromNumber(number int32) *OrderType_OrderTypeEnums {
	e := OrderType_OrderTypeEnums(number)
	return &e
}

func NewOrderType_OrderTypeEnumsFromValue(value string) *OrderType_OrderTypeEnums {
	return NewOrderType_OrderTypeEnumsFromNumber(OrderType_OrderTypeEnums_value[value])
}

func (e OrderType_OrderTypeEnums) GetDesc() string {
	return OrderType_OrderTypeEnums_name_desc[int32(e)]
}

func (OrderType_OrderTypeEnums) GetDescFromNumber(number int32) string {
	return OrderType_OrderTypeEnums_name_desc[number]
}

func (OrderType_OrderTypeEnums) GetDescFromName(name string) string {
	return OrderType_OrderTypeEnums_value_desc[name]
}

var (
	VoucherState_UseStateEnums_name_desc = map[int32]string{
		0: "领取状态",
		1: "使用状态",
		2: "过期状态",
	}

	VoucherState_UseStateEnums_value_desc = map[string]string{
		"NEW":     "领取状态",
		"USED":    "使用状态",
		"EXPIRED": "过期状态",
	}
)

func NewVoucherState_UseStateEnumsFromNumber(number int32) *VoucherState_UseStateEnums {
	e := VoucherState_UseStateEnums(number)
	return &e
}

func NewVoucherState_UseStateEnumsFromValue(value string) *VoucherState_UseStateEnums {
	return NewVoucherState_UseStateEnumsFromNumber(VoucherState_UseStateEnums_value[value])
}

func (e VoucherState_UseStateEnums) GetDesc() string {
	return VoucherState_UseStateEnums_name_desc[int32(e)]
}

func (VoucherState_UseStateEnums) GetDescFromNumber(number int32) string {
	return VoucherState_UseStateEnums_name_desc[number]
}

func (VoucherState_UseStateEnums) GetDescFromName(name string) string {
	return VoucherState_UseStateEnums_value_desc[name]
}

var (
	VoucherState_CategoryEnums_name_desc = map[int32]string{
		0: "注册赠券",
		1: "任务卷",
		2: "促销劵",
	}

	VoucherState_CategoryEnums_value_desc = map[string]string{
		"NEW_USER":  "注册赠券",
		"TASK":      "任务卷",
		"PROMOTION": "促销劵",
	}
)

func NewVoucherState_CategoryEnumsFromNumber(number int32) *VoucherState_CategoryEnums {
	e := VoucherState_CategoryEnums(number)
	return &e
}

func NewVoucherState_CategoryEnumsFromValue(value string) *VoucherState_CategoryEnums {
	return NewVoucherState_CategoryEnumsFromNumber(VoucherState_CategoryEnums_value[value])
}

func (e VoucherState_CategoryEnums) GetDesc() string {
	return VoucherState_CategoryEnums_name_desc[int32(e)]
}

func (VoucherState_CategoryEnums) GetDescFromNumber(number int32) string {
	return VoucherState_CategoryEnums_name_desc[number]
}

func (VoucherState_CategoryEnums) GetDescFromName(name string) string {
	return VoucherState_CategoryEnums_value_desc[name]
}

var (
	SmsType_SmsTypeEnums_name_desc = map[int32]string{
		0: "DELETED",
	}

	SmsType_SmsTypeEnums_value_desc = map[string]string{
		"TencentSms": "DELETED",
	}
)

func NewSmsType_SmsTypeEnumsFromNumber(number int32) *SmsType_SmsTypeEnums {
	e := SmsType_SmsTypeEnums(number)
	return &e
}

func NewSmsType_SmsTypeEnumsFromValue(value string) *SmsType_SmsTypeEnums {
	return NewSmsType_SmsTypeEnumsFromNumber(SmsType_SmsTypeEnums_value[value])
}

func (e SmsType_SmsTypeEnums) GetDesc() string {
	return SmsType_SmsTypeEnums_name_desc[int32(e)]
}

func (SmsType_SmsTypeEnums) GetDescFromNumber(number int32) string {
	return SmsType_SmsTypeEnums_name_desc[number]
}

func (SmsType_SmsTypeEnums) GetDescFromName(name string) string {
	return SmsType_SmsTypeEnums_value_desc[name]
}

var (
	PayStrategy_PayType_name_desc = map[int32]string{
		0: "支付宝",
		1: "微信",
		3: "雨点",
	}

	PayStrategy_PayType_value_desc = map[string]string{
		"ALI":      "支付宝",
		"WECHAT":   "微信",
		"INTEGRAL": "雨点",
	}
)

func NewPayStrategy_PayTypeFromNumber(number int32) *PayStrategy_PayType {
	e := PayStrategy_PayType(number)
	return &e
}

func NewPayStrategy_PayTypeFromValue(value string) *PayStrategy_PayType {
	return NewPayStrategy_PayTypeFromNumber(PayStrategy_PayType_value[value])
}

func (e PayStrategy_PayType) GetDesc() string {
	return PayStrategy_PayType_name_desc[int32(e)]
}

func (PayStrategy_PayType) GetDescFromNumber(number int32) string {
	return PayStrategy_PayType_name_desc[number]
}

func (PayStrategy_PayType) GetDescFromName(name string) string {
	return PayStrategy_PayType_value_desc[name]
}

var (
	CacheAdapterEnums_name_desc = map[int32]string{
		0: "REDIS",
	}

	CacheAdapterEnums_value_desc = map[string]string{
		"REDIS": "REDIS",
	}
)

func NewCacheAdapterEnumsFromNumber(number int32) *CacheAdapterEnums {
	e := CacheAdapterEnums(number)
	return &e
}

func NewCacheAdapterEnumsFromValue(value string) *CacheAdapterEnums {
	return NewCacheAdapterEnumsFromNumber(CacheAdapterEnums_value[value])
}

func (e CacheAdapterEnums) GetDesc() string {
	return CacheAdapterEnums_name_desc[int32(e)]
}

func (CacheAdapterEnums) GetDescFromNumber(number int32) string {
	return CacheAdapterEnums_name_desc[number]
}

func (CacheAdapterEnums) GetDescFromName(name string) string {
	return CacheAdapterEnums_value_desc[name]
}

var (
	EventMessageType_name_desc = map[int32]string{
		0:  "订单添加",
		1:  "延迟扣减用户金额",
		2:  "创建实例",
		3:  "关闭实例",
		4:  "开机",
		5:  "实例释放",
		20: "雨点赠送",
	}

	EventMessageType_value_desc = map[string]string{
		"ORDER_ADD":                "订单添加",
		"ORDER_EXPENSE_DEDUCTIONS": "延迟扣减用户金额",
		"ORDER_CREATE_INSEANCE":    "创建实例",
		"DEVICE_STOP_INSEANCE":     "关闭实例",
		"DEVICE_START_INSEANCE":    "开机",
		"DEVICE_RELEASE_INSEANCE":  "实例释放",
		"ACCOUNT_GRANT_INTEGRAL":   "雨点赠送",
	}
)

func NewEventMessageTypeFromNumber(number int32) *EventMessageType {
	e := EventMessageType(number)
	return &e
}

func NewEventMessageTypeFromValue(value string) *EventMessageType {
	return NewEventMessageTypeFromNumber(EventMessageType_value[value])
}

func (e EventMessageType) GetDesc() string {
	return EventMessageType_name_desc[int32(e)]
}

func (EventMessageType) GetDescFromNumber(number int32) string {
	return EventMessageType_name_desc[number]
}

func (EventMessageType) GetDescFromName(name string) string {
	return EventMessageType_value_desc[name]
}

var (
	OrderPyStateEnums_name_desc = map[int32]string{
		0: "未支付订单",
		1: "已经支付订单",
		2: "超时取消订单",
		3: "创建失败",
	}

	OrderPyStateEnums_value_desc = map[string]string{
		"NEW":    "未支付订单",
		"PAY":    "已经支付订单",
		"CANCEL": "超时取消订单",
		"FAIL":   "创建失败",
	}
)

func NewOrderPyStateEnumsFromNumber(number int32) *OrderPyStateEnums {
	e := OrderPyStateEnums(number)
	return &e
}

func NewOrderPyStateEnumsFromValue(value string) *OrderPyStateEnums {
	return NewOrderPyStateEnumsFromNumber(OrderPyStateEnums_value[value])
}

func (e OrderPyStateEnums) GetDesc() string {
	return OrderPyStateEnums_name_desc[int32(e)]
}

func (OrderPyStateEnums) GetDescFromNumber(number int32) string {
	return OrderPyStateEnums_name_desc[number]
}

func (OrderPyStateEnums) GetDescFromName(name string) string {
	return OrderPyStateEnums_value_desc[name]
}

var (
	InstanceStatusEnums_name_desc = map[int32]string{
		0: "开机",
		1: "关机",
		2: "失败",
		3: "初始化",
		4: "异常",
		5: "重新创建",
	}

	InstanceStatusEnums_value_desc = map[string]string{
		"On":        "开机",
		"Off":       "关机",
		"Fail":      "失败",
		"Invalid":   "初始化",
		"Exception": "异常",
		"ReCreate":  "重新创建",
	}
)

func NewInstanceStatusEnumsFromNumber(number int32) *InstanceStatusEnums {
	e := InstanceStatusEnums(number)
	return &e
}

func NewInstanceStatusEnumsFromValue(value string) *InstanceStatusEnums {
	return NewInstanceStatusEnumsFromNumber(InstanceStatusEnums_value[value])
}

func (e InstanceStatusEnums) GetDesc() string {
	return InstanceStatusEnums_name_desc[int32(e)]
}

func (InstanceStatusEnums) GetDescFromNumber(number int32) string {
	return InstanceStatusEnums_name_desc[number]
}

func (InstanceStatusEnums) GetDescFromName(name string) string {
	return InstanceStatusEnums_value_desc[name]
}

var (
	BillingTypeEnums_name_desc = map[int32]string{
		0: "按量",
		1: "按日",
		2: "按月",
		3: "周",
		4: "年",
		5: "续费",
	}

	BillingTypeEnums_value_desc = map[string]string{
		"AMOUNT": "按量",
		"DAY":    "按日",
		"MONTH":  "按月",
		"WEEK":   "周",
		"YEAR":   "年",
		"RENEW":  "续费",
	}
)

func NewBillingTypeEnumsFromNumber(number int32) *BillingTypeEnums {
	e := BillingTypeEnums(number)
	return &e
}

func NewBillingTypeEnumsFromValue(value string) *BillingTypeEnums {
	return NewBillingTypeEnumsFromNumber(BillingTypeEnums_value[value])
}

func (e BillingTypeEnums) GetDesc() string {
	return BillingTypeEnums_name_desc[int32(e)]
}

func (BillingTypeEnums) GetDescFromNumber(number int32) string {
	return BillingTypeEnums_name_desc[number]
}

func (BillingTypeEnums) GetDescFromName(name string) string {
	return BillingTypeEnums_value_desc[name]
}

var (
	InvoiceType_name_desc = map[int32]string{
		0: "增值税普通发票",
		1: "增值税专用发票",
	}

	InvoiceType_value_desc = map[string]string{
		"VAT_PLAIN_INVOICE":   "增值税普通发票",
		"VAT_SPECIAL_INVOICE": "增值税专用发票",
	}
)

func NewInvoiceTypeFromNumber(number int32) *InvoiceType {
	e := InvoiceType(number)
	return &e
}

func NewInvoiceTypeFromValue(value string) *InvoiceType {
	return NewInvoiceTypeFromNumber(InvoiceType_value[value])
}

func (e InvoiceType) GetDesc() string {
	return InvoiceType_name_desc[int32(e)]
}

func (InvoiceType) GetDescFromNumber(number int32) string {
	return InvoiceType_name_desc[number]
}

func (InvoiceType) GetDescFromName(name string) string {
	return InvoiceType_value_desc[name]
}

var (
	VoucherPublishState_name_desc = map[int32]string{
		0: "发布",
		1: "草稿",
		2: "下线",
	}

	VoucherPublishState_value_desc = map[string]string{
		"PUBLISH": "发布",
		"DRAFT":   "草稿",
		"OFFLINE": "下线",
	}
)

func NewVoucherPublishStateFromNumber(number int32) *VoucherPublishState {
	e := VoucherPublishState(number)
	return &e
}

func NewVoucherPublishStateFromValue(value string) *VoucherPublishState {
	return NewVoucherPublishStateFromNumber(VoucherPublishState_value[value])
}

func (e VoucherPublishState) GetDesc() string {
	return VoucherPublishState_name_desc[int32(e)]
}

func (VoucherPublishState) GetDescFromNumber(number int32) string {
	return VoucherPublishState_name_desc[number]
}

func (VoucherPublishState) GetDescFromName(name string) string {
	return VoucherPublishState_value_desc[name]
}
