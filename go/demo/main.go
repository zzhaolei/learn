package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var client = http.DefaultClient

var cookies = ""

var url = "https://shell.obrase.com/hpiApp/web-isure/log/doQueryStockOutOperateLogWithCost.do?operateTimeStart=2023-02-17%2000%3A00%3A00&operateTimeEnd=2023-05-17%2023%3A59%3A59&tid=&outerSkuId=&skuColor=&skuSize=&goodsNo=&userName=&operateName=%E4%BE%BF%E6%8D%B7%E5%BA%93%E5%AD%98%E5%87%BA%E5%BA%93&stockOperateType=&stockOutNumQuery=&groupBy=&depotName=&rows=100&page=1&queryCost=1&_=1684303068308"

func main() {
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Fatalln("1: 创建 request 对象失败")
	}

	// SetCookie
	for _, cookie := range strings.Split(cookies, "; ") {
		cookie := strings.Split(cookie, "=")
		req.AddCookie(&http.Cookie{
			Name:  cookie[0],
			Value: cookie[1],
		})
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("2: 发送请求失败 s")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
