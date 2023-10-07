package response

import (
	"github.com/go-chi/render"
	"net/http"
)

//
//type Response struct {
//	Status string `json:"status"`
//	Error  string `json:"error,omitempty"`
//}

type Response struct {
	Err            error    `json:"-"` // низкоуровневая ошибка исполнения
	HTTPStatusCode int      `json:"-"` // HTTP статус код
	ErrorMessage   *Details `json:"error"`
}

type Details struct {
	StatusText  string `json:"status"`            // сообщение пользовательского уровня
	MessageText string `json:"message,omitempty"` // application-level сообщение, для дебага
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func InternalError(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessage: &Details{
			AppCode:     http.StatusInternalServerError,
			StatusText:  "Internal Server Error",
			MessageText: err.Error(),
		},
	}
}
