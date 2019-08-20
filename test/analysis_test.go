package test

import (
	"GoProxyPool/defs"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"strings"
	"testing"
)

func TestAnalysis(t *testing.T) {
	url := "http://www.66ip.cn/areaindex_1/1.html"
	data, _ := httplib.EuUserGetEncoding(url)
	html := string(data)
	document, e := goquery.NewDocumentFromReader(strings.NewReader(html))
	if e != nil {
		clog.Println(e.Error())
		return
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
			datalist[i-1].Port = text
		}
	})

	//reg := `<tr>([\d\D]+?)<tr>`
	//
	//compile := regexp.MustCompile(reg)
	//submatch := compile.FindAllStringSubmatch(html, -1)
	//
	//for _,k := range submatch {
	//	data := k[1]
	//
	//	reg := `<td>([\D\d]+?)<td>`
	//	compile := regexp.MustCompile(reg)
	//	submatch := compile.FindAllStringSubmatch(data, -1)
	//
	//	for _,k :=range submatch {
	//		log.Println(k[0])
	//	}
	//}

}
