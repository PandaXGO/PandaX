package model

// 分页参数
type PageParam struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Params   string `json:"params"`
}

type ResultPage struct {
	Total    int64 `json:"total"`
	PageNum  int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
	Data     any   `json:"data"`
}
