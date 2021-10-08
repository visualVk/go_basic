package parser

import (
	"example.com/ch4/src/crawler/engine"
	"fmt"
	"regexp"
)

const REGEXP = `<a[^href]*href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParserCityList(content []byte) engine.ParserResult {
	re := regexp.MustCompile(REGEXP)
	matches := re.FindAllSubmatch(content, -1)
	var parserResult engine.ParserResult

	for _, m := range matches {
		fmt.Printf("City: %s, url: %s\n", m[2], m[1])
		var request = engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParserFunc,
		}

		parserResult.Requests = append(parserResult.Requests, request)
		parserResult.Items = append(parserResult.Items, string(m[2]))
	}

	return parserResult

}
