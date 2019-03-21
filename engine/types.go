package engine

type Request struct {
	Url string
	ParserFunc
}

type ParserFunc func(content []byte, url string) ParseResult

type ParseResult struct {
	Requests []Request
	Items    []Item
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Id      string
	Type    string
	PayLoad interface{}
}
