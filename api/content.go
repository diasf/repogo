package api

type ContentService interface {
	GetContents(request GetContentsRequest) GetContentsResponse
}

type GetContentsRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type ContentsResponse struct {
	Total  int       `json:"total"`
	Result []Content `json:"result"`
}

type Content struct {
	ID     string `json:"id"`
	Parent string `json:"parent"`
	Label  string `json:"label"`
}
