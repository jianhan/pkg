package http

type Validator interface {
	Validate() error
}
