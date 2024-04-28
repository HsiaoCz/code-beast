package handlers

type APIError struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

func (a APIError) Error() string {
	return a.Msg
}
