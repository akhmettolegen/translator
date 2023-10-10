package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
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
