package excel_model

type ExportDailyStatementExcel struct {
	Date              string `excel:"name:日期;"`     // 日期
	SerialNumber      string `excel:"name:流水号;"`    // 流水号
	BillCreateTime    string `excel:"name:账单创建时间;"` // 账单创建时间
	TransactionType   string `excel:"name:交易类型;"`   // 消费类型
	TransactionAmount string `excel:"name:交易金额;"`   // 交易金额
	DiscountAmount    string `excel:"name:优惠金额;"`   // 优惠金额
	BalancePayment    string `excel:"name:余额支付;"`   // 余额支付
	VoucherDeduction  string `excel:"name:代金券抵扣;"`  // 代金券抵扣
}
