package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

	const cityListRe = `<a[\s]+href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte, _ string) engine.ParseResult {
	//<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range all {
		//result.Items = append(result.Items, "City:"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	log.Printf("总共获取城市有:%d", len(all))
	return result
}
