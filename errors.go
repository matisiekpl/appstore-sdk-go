package appstore_sdk

type ErrorResult struct {
	Errors []*Error `json:"errors"`
}

func (e *ErrorResult) GetError() *Error {
	return e.Errors[0]
}

type Error struct {
	Id     string       `json:"id"`
	Status string       `json:"status"`
	Code   string       `json:"code"`
	Title  string       `json:"title"`
	Detail string       `json:"detail"`
	Source *ErrorSource `json:"source"`
}

type ErrorSource struct {
	Parameter string `json:"parameter"`
	Pointer   string `json:"pointer"`
}
