package serviceErrors

import "venturan/app/common/response"

// Mapping以1040000开头
var (
	CommodityStatusIsNotEmpty = response.Response{1040000, nil, "商品状态不能为空"}
)
