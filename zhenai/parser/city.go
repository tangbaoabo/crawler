package parser

import (
	"crawler/engine"
	"regexp"
)

var reg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var linkRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

//<a href="http://www.zhenai.com/zhenghun/shanghai/2">下一页</a>
//<a target="_blank" href="http://www.zhenai.com/zhenghun/xuhui">徐汇征婚</a>
//<a href="http://album.zhenai.com/u/1414260922" target="_blank">暮荣</a>
//<td width="180"><span class="grayL">性别：</span>男士</td>

func ParseCity(contents []byte, _ string) engine.ParseResult {
	allMatch := reg.FindAllSubmatch(contents, -1)
	parseResult := engine.ParseResult{}
	sex := sexRe.FindAllSubmatch(contents, -1)
	//for _, value := range allMatch {
	//	parseResult.Items = append(parseResult.Items, "User:"+string(value[2]))
	//	parseResult.Requests = append(parseResult.Requests,
	//		engine.Request{
	//			Url: string(value[1]),
	//			ParserFunc: func(bytes []byte) engine.ParseResult {
	//				return ParseUser(bytes, map[string]interface{}{"sex": string(sex[1])})
	//			},
	//		})
	//}

	for i := 0; i < len(allMatch); i++ {
		//parseResult.Items = append(parseResult.Items, "User:"+string(allMatch[i][2]))
		parseResult.Requests = append(parseResult.Requests,
			engine.Request{
				Url:        string(allMatch[i][1]),
				ParserFunc: FileParse(string(sex[i][1])),
			})
	}
	//log.Printf("总共获取城市有:%d", len(allMatch))

	matches := linkRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		parseResult.Requests = append(parseResult.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return parseResult
}

func FileParse(sex string) engine.ParserFunc {
	return func(content []byte, url string) engine.ParseResult {
		return ParseUser(content, map[string]interface{}{
			"sex": sex,
			"url": url,
		})
	}
}
