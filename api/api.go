package api

import (
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_ERROR = "error"
)

type ResponseSuccessData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseSuccessLoginData struct {
	ResponseSuccessData
	Token string `json:"token"`
}

type ResponseErrorData struct {
	ResponseSuccessData
	Message string `json:"message"`
}

type RespondMessage struct {
	Message string
}

type Option func(*RespondMessage)

func WithMessageError(message string) Option {
	return func(r *RespondMessage) {
		r.Message = message
	}
}

func RespondSuccess(w *gin.Context, code int, obj interface{}) {

	var res ResponseSuccessData

	res.Status = code
	res.Data = obj

	w.JSON(code, res)
}

func RespondLoginSuccess(w *gin.Context, code int, obj interface{}, token string) {

	var res ResponseSuccessLoginData

	res.Status = code
	res.Data = obj
	res.Token = token

	w.JSON(code, res)
}

func RespondError(w *gin.Context, code int, option ...Option) {
	var res ResponseErrorData
	r := &RespondMessage{}
	res.Status = code
	for _, o := range option {
		o(r)
	}

	if r.Message == "" {
		res.Message = DEFAULT_ERROR
	} else {
		res.Message = r.Message
	}

	res.Data = map[string]string{} // return {}
	// res.Data = []string{} // return []

	w.JSON(code, res)
}
