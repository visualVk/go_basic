package engine

import (
	"example.com/ch4/src/crawler/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request){
	var requests []Request

	for _, seed := range seeds{
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		log.Printf("fetching url: %s\n", request.Url)
		content, err := fetcher.Fetch(request.Url)

		if err != nil {
			log.Printf("fetch error: %v", err)
			continue
		}

		cityListParse := request.ParserFunc(content)
		for _, cityName := range cityListParse.Items{
			fmt.Printf("city: %v\n", cityName)
		}

		requests = append(requests, cityListParse.Requests...)
	}
}
