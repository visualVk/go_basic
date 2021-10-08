package main

import (
	"example.com/ch4/src/crawler/engine"
	"example.com/ch4/src/crawler/parser/citylist"
)

func main() {

	// err = os.WriteFile("src/crawler/output.out", content, os.ModeAppend.Perm())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", content)
	//content, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}

	//parserResult := parser.ParserCityList(content)
	//fmt.Printf("%#v", parserResult)

	engine.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
