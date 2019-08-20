package reptile

import (
	"GoProxyPool/defs"
	"GoProxyPool/detasource/memory_conn"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const testurl = "http://www.bing.com"

func reptileAbroad() {
	log.Println("1---")
	const wat = "http://www.66ip.cn/1.html"
	data := defs.ProxyList{}

	// 爬取前10页
	for i := 1; i <= 10; i++ {
		url := "http://www.66ip.cn/" + strconv.Itoa(i) + ".html"

		bytes, e := httplib.EuUserGetEncoding(url)

		if e == nil {
			dalist, err := analysis(bytes)
			if err == nil {
				data = checkProxy(dalist)
			}
		}
	}

	bytes, e := json.Marshal(data)
	if e == nil {
		memory_conn.MemoryDb.Store("gw", bytes)
	}
}

func reptileDomestic() {
	const dom = "http://www.66ip.cn/areaindex_1/1.html" // 1~10
	data := defs.ProxyList{}

	// 爬取前10页
	for i := 1; i <= 10; i++ {
		url := "http://www.66ip.cn/areaindex_" + strconv.Itoa(i) + "/1.html"

		bytes, e := httplib.EuUserGetEncoding(url)

		if e == nil {
			dalist, err := analysis(bytes)
			if err == nil {
				data = checkProxy(dalist)
			}
		}
	}

	bytes, e := json.Marshal(data)
	if e == nil {
		memory_conn.MemoryDb.Store("gn", bytes)
	}
}

// 分析页面
func analysis(data []byte) (defs.ProxyList, error) {
	html := string(data)
	document, e := goquery.NewDocumentFromReader(strings.NewReader(html))
	if e != nil {
		clog.Println(e.Error())
		return nil, e
	}
	datalist := defs.ProxyList{}

	tbody := document.Find("tbody")
	tbody.Find("td:nth-child(1)").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text != "ip" {
			data := defs.Proxy{Ip: text}
			datalist = append(datalist, &data)
		}
	})

	tbody.Find("td:nth-child(2)").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text != "端口号" {
			datalist[i-1].Port = text
		}
	})

	tbody.Find("td:nth-child(3)").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text != "代理位置" {
			datalist[i-1].Area = text
		}
	})

	return datalist, nil
}

// 验证代理是否可用
func checkProxy(list defs.ProxyList) defs.ProxyList {
	for k, v := range list {
		request := httplib.Get(testurl).SetProxy(func(request *http.Request) (*url.URL, error) {
			u := new(url.URL)
			u.Scheme = "http"
			u.Host = v.Ip + ":" + v.Port //蓝灯的http代理地址
			return u, nil
		})

		request.SetTimeout(time.Second*5, time.Second*5) //设置超时时间

		_, e := request.String()
		if e != nil {
			if len(list) != 0 {
				// 如果验证失败 则删除
				clog.Println(e.Error())
				log.Println("失效   ====" + v.Ip + v.Port)
				if k+1<len(list)  {
					list = append(list[:k], list[k+1:]...)
				}else {
					list = list[:k]
				}
			}

		}else {
			bytes, e := json.Marshal(v)

			if e == nil {
				memory_conn.MemoryDb.Store(1,string(bytes))
			}

			log.Println("成功   ====")
		}
	}

	return list
}
