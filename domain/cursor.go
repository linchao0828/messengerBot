package domain

const (
	defaultCursorLimit = 10
)

type Cursor struct {
	CursorID   int64 `json:"cursor_id" form:"cursor_id" example:"0"`      //翻页游标，为0时代表取最新的一页，大于0时代表取cursor_id之前的一页
	CursorSize int   `json:"cursor_size" form:"cursor_size" example:"10"` //游标翻页每页条数
}

func (req *Cursor) LastID() int64 {
	if req.CursorID > 0 {
		return req.CursorID
	}
	return 0
}

func (req *Cursor) Limit() int {
	if req.CursorSize > 0 {
		return req.CursorSize
	}
	return defaultCursorLimit
}
