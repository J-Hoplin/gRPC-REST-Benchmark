package service

type CommonQuery struct {
	From int64 `form:"from" binding:"required,gte=0"`
	To   int64 `form:"to" binding:"required,gte=0"`
}

type RestResponse struct {
	ResponseNumber int `json:"response_number"`
}
