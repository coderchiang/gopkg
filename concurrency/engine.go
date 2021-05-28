package concurrency

import (
	"pkg/concurrency/scheduler"
)

type Engine struct {
	Scheduler   scheduler.Scheduler
	WorkerCount int
}

// 并发版
func (e *Engine) Run(seeds ...interface{}) {
	out := make(chan interface{})
	e.Scheduler.Run()
    //开启并发work
	for i := 0; i < e.WorkerCount; i++ {
		go e.createWorker(e.Scheduler.NewWorkerChan(), out, e.Scheduler)
	}
    //提交并发seed
	for _, r := range seeds {
		//todo
		e.Scheduler.Submit(r)
	}
	//处理运行后的结果
	for result :=range out{
		//todo
		println(result)
	}


}


func (e Engine) createWorker(in chan interface{}, out chan interface{}, ready scheduler.ReadyNotifier) {
		for {
			// Tell scheduler I am ready
			ready.WorkerChanReady(in)
			res := <-in
			result, err := e.doWorker(res)
			if err != nil {
				continue
			}
			out <- result
		}
}


func (Engine) doWorker(res  interface{}) (interface{},error) {

	//todo
	return res,nil
}

