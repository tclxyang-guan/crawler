package parse

import (
	"io/ioutil"
	"log"
	"net/http"
	"pc/models"
	"regexp"
)

func Get(Url string) ([]byte, error) {
	resp, err := http.Get(Url)
	if err != nil {
		log.Println("zhenai data return fail:", err)
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

const CityRegexp = `<a[^href]*[^>]href="(http://www.zhenai.com/zhenghun/[a-z]+)"[^>]*>([^<]+)</a>`
const UserRegexp = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const NextUserRegexp = `<a href="http://www.zhenai.com/zhenghun/[^"]+">下一页</a>`

func ParseCity(b []byte) (pr models.ParseResult) {
	city := regexp.MustCompile(CityRegexp)
	str := city.FindAllStringSubmatch(string(b), -1)
	for _, row := range str {
		pr.Requests = append(pr.Requests, models.Request{
			Url:       string(row[1]),
			ParseFunc: ParseUser,
		})
		pr.Data = append(pr.Data, row[2])
	}
	return
}

func ParseUser(b []byte) (pr models.ParseResult) {
	city := regexp.MustCompile(UserRegexp)
	str := city.FindAllStringSubmatch(string(b), -1)
	for _, row := range str {
		pr.Requests = append(pr.Requests, models.Request{
			Url:       string(row[1]),
			ParseFunc: models.NewParseFunc,
		})
		pr.Data = append(pr.Data, row[2])
	}
	return
}
