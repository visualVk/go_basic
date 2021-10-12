package main

import (
	"example.com/ch4/src/crawler/engine"
	"example.com/ch4/src/crawler/parser/citylist"
	"example.com/ch4/src/crawler/scheduler"
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
	e := engine.ConcurrentEngine{Scheduler: &scheduler.Scheduler{}, WorkerCount: 100}

	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
