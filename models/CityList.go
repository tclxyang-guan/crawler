package models

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}
type ParseResult struct {
	Requests []Request
	Data     []interface{}
}

func NewParseFunc([]byte) ParseResult {
	return ParseResult{}
}
