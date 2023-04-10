package model

import "github.com/cassa10/arq2-tp1/src/domain/util"

type PagingRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type Paging struct {
	Total       int `json:"total"`
	PageSize    int `json:"pageSize"`
	Pages       int `json:"pages"`
	CurrentPage int `json:"currentPage"`
}

func (pr PagingRequest) String() string {
	return util.ParseStruct("PagingRequest", pr)
}

func (p Paging) String() string {
	return util.ParseStruct("Paging", p)
}

func NewPaging(total int, pageSize int, pages int, currentPage int) Paging {
	return Paging{Total: total, PageSize: pageSize, Pages: pages, CurrentPage: currentPage}
}

func EmptyPage() Paging {
	return NewPaging(0, 0, 0, 0)
}
