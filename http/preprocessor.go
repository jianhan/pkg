package http

type Preprocessor interface {
	Preprocess() error
}
