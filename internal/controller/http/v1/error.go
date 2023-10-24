package v1

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	ErrMessage     string `json:"error"`
	HTTPStatusCode int    `json:"-"`
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func errResponse(code int, msg string) render.Renderer {
	return &Response{
		ErrMessage:     msg,
		HTTPStatusCode: code,
	}
}
