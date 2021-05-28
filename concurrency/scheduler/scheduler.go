package scheduler


type Scheduler interface {
	Submit(interface{})

	NewWorkerChan() chan interface{}

	ReadyNotifier

	Run()
}

type ReadyNotifier interface {
	WorkerChanReady(chan interface{})
}
