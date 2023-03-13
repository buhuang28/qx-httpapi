package dto

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) Result {
	return Result{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func Fail(msg string) Result {
	return Result{
		Code:    -1,
		Message: msg,
	}
}
