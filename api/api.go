package api

import (
	"github.com/gin-gonic/gin"
)

type ResponseSuccessData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseSuccessLoginData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}

type ResponseErrorData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

func RespondError(w *gin.Context, code int, obj string) {
	var res ResponseErrorData

	res.Status = code
	res.Message = obj
	res.Data = map[string]string{} // return {}
	// res.Data = []string{} // return []

	w.JSON(code, res)
}
