package main

import (
	"fmt"
	"pc/models"
	"pc/parse"
)

func OneRun(request ...models.Request) {
	var reqs []models.Request
	for _, v := range request {
		reqs = append(reqs, v)
	}
	for len(reqs) > 0 {
		req := reqs[0]
		reqs = reqs[1:]
		fmt.Println("正在调取Url:", req.Url)
		b, err := parse.Get(req.Url)
		if err != nil {
			continue
		}
		pr := req.ParseFunc(b)
		reqs = append(reqs, pr.Requests...)
		for _, v := range pr.Data {
			fmt.Println("值为:", v)
		}
	}
}
func main() {
	OneRun(models.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parse.ParseCity,
	})

}
