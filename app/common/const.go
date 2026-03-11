package common

const (
	InactiveStatus int = iota // 禁用
	ActionStatus              // 正常
)

const (
	NotDeleted int8 = iota // 未删除
	Deleted    int8 = 1    // 已删除
)

// 商品状态
const (
	NotAvailable int = iota // 未上架
	Available               // 在售
	Removed                 // 下架
	SaleOut                 // 售罄
)
