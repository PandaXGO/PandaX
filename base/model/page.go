package model

// 分页参数
type PageParam struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Params   string `json:"params"`
}

// 分页结果
type PageResult struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}
