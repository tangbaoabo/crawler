package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan Item
}

type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}
type ReadNotifier interface {
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	itemCount := 1
	//flag := true
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("the result is #%d: %v", itemCount, item)
			itemCount++
			go func() {
				e.ItemChan <- item
			}()
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
		////处理循环计数
		//if flag {
		//	itemCount = 1
		//	flag = false
		//}
	}
}

var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false

}

func createWorker(in chan Request, out chan<- ParseResult, notifier ReadNotifier) {
	go func() {
		for {
			//tell scheduler i'm ready
			notifier.WorkReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
