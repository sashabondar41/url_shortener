package dto

type GetLongRequest struct {
	Short string `json:"short"`
}

type GetLongResponse struct {
	Long string `json:"long"`
}
