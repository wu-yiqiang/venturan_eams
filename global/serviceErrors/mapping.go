package serviceErrors

import "venturan/app/common/response"

// Mapping以1030000开头
var (
	MappingTypeIsNotEmpty = response.Response{1030000, nil, "字典类型不能为空"}
)
