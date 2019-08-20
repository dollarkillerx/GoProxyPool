package test

import (
	"github.com/dollarkillerx/easyutils/httplib"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const testurl = "https://www.bing.com"

func TestCheckProxy(t *testing.T) {
	request := httplib.Get(testurl).SetProxy(func(request *http.Request) (*url.URL, error) {
		u := new(url.URL)
		u.Scheme = "http"
		u.Host = "101.108.9.171:8080" //蓝灯的http代理地址
		return u, nil
	})

	request.SetTimeout(time.Second*5, time.Second*5) //设置超时时间

	s, e := request.String()
	if e == nil {
		log.Println(s)
	} else {
		panic(e.Error())
	}
}
