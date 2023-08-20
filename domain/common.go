package domain

type File struct {
	Key string `json:"key"` // 文件后缀
	Url string `json:"url"` // 文件地址
}
