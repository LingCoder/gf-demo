package base

import "gf-demo/library/consts"

type IPage interface {
	GetPageNum() int
	GetPageSize() int
	GetOrders() []OrderItem
}

type Page struct {
	PageNum  int         `json:"pageNum" form:"pageNum,default=1"`
	PageSize int         `json:"pageSize" form:"pageSize,default=10"`
	Orders   []OrderItem `json:"orders" form:"orders"`
}

func (p Page) GetPageNum() int {
	return p.PageNum
}

func (p Page) GetPageSize() int {
	return p.PageSize
}

func (p Page) GetOrders() []OrderItem {
	return p.Orders
}

type OrderItem struct {
	Column string
	Asc    bool
}

func (s *PageReq) PageInfo() (offset int, limit int) {
	if s.PageNum > 0 && s.PageSize > 0 {
		offset = int((s.PageNum - 1) * s.PageSize)
		limit = int(s.PageSize)
	} else {
		offset = 0
		limit = consts.DefaultPageSize
	}
	return
}

type IPageInfo[T any] interface {
	GetTotal() int
	GetList() []*T
	SetPageInfo(IPageInfo[T])
}

type PageInfo[T any] struct {
	List  []*T
	Total int
}

func (p *PageInfo[T]) SetPageInfo(i IPageInfo[T]) {
	if i != nil {
		p.Total = i.GetTotal()
		p.List = i.GetList()
	}
}

func (p *PageInfo[T]) GetTotal() int {
	return p.Total
}

func (p *PageInfo[T]) GetList() []*T {
	return p.List
}
