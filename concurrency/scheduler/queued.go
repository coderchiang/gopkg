package scheduler


// Request队列和Worker队列
type QueuedScheduler struct {
	inChan chan interface{}
	//双层chan逻辑
	workChans    chan chan interface{}
}

//提交信息进入inChan
func (s *QueuedScheduler) Submit(r interface{}) {
	s.inChan <- r
}
//初始化工作的inChan
func (s *QueuedScheduler) NewWorkerChan() chan interface{} {
	return make(chan interface{})
}
//inChan加入到workChans
func (s *QueuedScheduler) WorkerChanReady(w chan interface{}) {
	s.workChans <- w
}

func (s *QueuedScheduler) Run() {
	s.workChans = make(chan chan interface{})
	s.inChan = make(chan interface{})
	go func() {
		var requestQ []interface{}
		var workQ []chan interface{}
		for {
			var activeRequest interface{}
			var activeWorker chan interface{}
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeWorker = workQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.inChan:
				requestQ = append(requestQ, r)
			case w := <-s.workChans:
				workQ = append(workQ, w)
			case activeWorker <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
