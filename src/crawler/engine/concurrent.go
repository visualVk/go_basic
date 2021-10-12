package engine

import (
	"example.com/ch4/src/crawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterChannel(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMasterChannel(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("%#v\n", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in

			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

func worker(request Request) (ParserResult, error) {
	log.Printf("fetching url: %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)

	if err != nil {
		log.Printf("fetch error: %v", err)
		return ParserResult{}, err
	}

	return request.ParserFunc(content), nil
}
