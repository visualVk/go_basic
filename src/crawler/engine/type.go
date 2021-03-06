package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []interface{}
}

func NilParserFunc([] byte) ParserResult{
	return ParserResult{}
}