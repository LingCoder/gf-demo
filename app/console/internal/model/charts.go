package model

type ChartsKlineInput struct {
}

type ChartsKlineOutput struct {
	List []*ChartsKlineItem
}

type ChartsKlineItem struct {
	BizDate string `json:"bizDate"`
	Amount  string `json:"amount"`
}
