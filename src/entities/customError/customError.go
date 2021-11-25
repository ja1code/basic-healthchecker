package customError

type HttpError struct {
	Url        string
	StatusCode int
	Body       []byte
}
