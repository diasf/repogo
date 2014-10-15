package api

func FindContents(param FindContentsParam) FindContentsResponse {

}

type FindContentsParam struct {
	Tags     []string
	Page     int32
	PageSize int32
}

type FindContentsResponse struct {
	Code200 ContentCollectionResponse
	Default ErrorResponse
}

type ErrorResponse struct {
	ErrorResource
}

type ContentCollectionResponse struct {
	CollectionResource
	Contents []ContentResource `json:"contents"`
}

type ContentResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CollectionResource struct {
	Meta MetaResource `json:"meta"`
}

type MetaResource struct {
	Total  int64 `json:"total"`
	Page   int64 `json:"page"`
	IsLast bool  `json:"isLast"`
}

type ErrorResource struct {
	Code    int64 `json:"code"`
	Message bool  `json:"message"`
}
