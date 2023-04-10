package dto

type PagingParamQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func (qs PagingParamQuery) GetSize() int {
	if qs.PageSize < 1 || qs.PageSize > 200 {
		return 10
	}
	return qs.PageSize
}

func (qs PagingParamQuery) GetPage() int {
	if qs.Page < 1 {
		return 0
	}
	if (qs.Page-1)*qs.GetSize() >= 10000 {
		return (10000 / qs.GetSize()) - 1
	}
	return qs.Page - 1
}
