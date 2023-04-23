package scheduler

import (
	"imooc.com/ccmouse/learngo/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
//	s.workerChan = c
//}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	go func() {
		s.workerChan <- r
	}()

}
