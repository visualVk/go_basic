package scheduler

import "example.com/ch4/src/crawler/engine"

type Scheduler struct {
	channel chan engine.Request
}

func (s *Scheduler) Submit(request engine.Request) {
	//submit should be processed by go routine each other
	//if we don't use goroutine, in channel will be filled with requests, while processing of parsing
	//request is behind with submit function, which finally will happen to deadlock
	go func() {
		s.channel <- request
	}()
}

func (s *Scheduler) ConfigureMasterChannel(requests chan engine.Request) {
	s.channel = requests
}
