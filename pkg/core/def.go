package core

const (
	InvalidParameter    = 1
	InternalServerError = -1
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
