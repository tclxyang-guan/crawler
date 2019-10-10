package main

import (
	"fmt"
	"pc/models"
	"pc/parse"
)

//将request给work
type Schedule struct {
	WorkChan chan models.Request
}

func (s *Schedule) SendWork(w models.Request) {
	go func() {
		s.WorkChan <- w
	}()

}
func (s *Schedule) InitWorkChan(wc chan models.Request) {
	s.WorkChan = wc
}

type Engine struct {
	Schedule  *Schedule
	WorkCount int
}

func createWork(in chan models.Request, out chan models.ParseResult) {
	go func() {
		for {
			result := <-in
			pr, err := worker(result)
			if err != nil {
				continue
			}
			out <- pr
		}
	}()
}
func worker(req models.Request) (models.ParseResult, error) {
	b, err := parse.Get(req.Url)
	if err != nil {
		return models.ParseResult{}, err
	}
	return req.ParseFunc(b), nil
}
func (s *Engine) MoreRun(request ...models.Request) {
	//var reqs []models.Request
	in := make(chan models.Request)
	out := make(chan models.ParseResult)
	//将schedule中的workChan初始化
	s.Schedule.InitWorkChan(in)
	//把所有的request都给workChan
	for _, v := range request {
		s.Schedule.SendWork(v)
	}
	for i := 0; i < s.WorkCount; i++ {
		createWork(in, out)
	}
	i := 0
	for {
		pr := <-out
		for _, v := range pr.Requests {
			s.Schedule.SendWork(v)
		}
		for _, v := range pr.Data {
			fmt.Println(i, "值为:", v)
			i++
		}

	}
}
func main() {

	e := Engine{
		Schedule:  &Schedule{},
		WorkCount: 100,
	}
	e.MoreRun(models.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parse.ParseCity,
	})
}
