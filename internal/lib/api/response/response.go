package response

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

func OK() Response {
	return Response{
		Status: 200,
	}
}

func Error(status int, msg string) Response {
	return Response{
		Status: status,
		Error:  msg,
	}
}
