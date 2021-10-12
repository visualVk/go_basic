package engine

import (
	"example.com/ch4/src/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		cityListParse, err := e.worker(request)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}

		requests = append(requests, cityListParse.Requests...)
	}
}

func (e SimpleEngine) worker(request Request) (ParserResult, error) {
	log.Printf("fetching url: %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)

	if err != nil {
		log.Printf("fetch error: %v", err)
		return ParserResult{}, err
	}

	return request.ParserFunc(content), nil
}
