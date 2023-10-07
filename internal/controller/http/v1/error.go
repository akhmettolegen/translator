package v1

import "context"

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c context.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}
