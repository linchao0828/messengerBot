package domain

const (
	defaultPageLimit = 10
	pageLimitPeriod  = 5000
)

type Pagination struct {
	PageNo   int `json:"page_no" form:"page_no" example:"1"`
	PageSize int `json:"page_size" form:"page_size" example:"10"`
}

func (req *Pagination) Offset() int {
	if no := req.PageNo - 1; no > 0 {
		return no * req.Limit()
	}
	return 0
}

func (req *Pagination) Limit() int {
	if req.PageSize > 0 && req.PageSize <= pageLimitPeriod {
		return req.PageSize
	}
	return defaultPageLimit
}
