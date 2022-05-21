package exceptions

type HTTPError struct {
	Code    string `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func (e HTTPError) Error() string {
	return e.Code
}

var ErrGameIsOver = HTTPError{
	Code:    "game_is_over",
	Message: "Game is over",
}
