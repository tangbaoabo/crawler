package engine

import (
	"crawler/fetcher"
	"log"
)

//将请求中的ur发送到fetcher，解析成下一组请求和内容
func worker(req Request) (ParseResult, error) {
	//log.Printf("Fetching %s", req.Url)
	//拉取网页中UTF8的文本内容
	content, err := fetcher.Fetch(req.Url)
	if err != nil {
		log.Printf("Fetcher :error fetching url %s :%v", req.Url, err)
		return ParseResult{}, nil
	}
	//解析拉取的内容
	return req.ParserFunc(content, req.Url), nil
}
