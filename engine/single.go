package engine

import (
	"log"
)
type Single struct {
}


func (single Single)Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		parseResult, e := worker(req)
		if e != nil {
			continue
		}
		//将新的解析的URL放入队列
		requests = append(requests, parseResult.Requests...)
		//生产内容
		for _, value := range parseResult.Items {
			log.Printf("the result is: %v", value)
		}
	}
}


