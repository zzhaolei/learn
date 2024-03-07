package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
)

type Type int

const (
	Title Type = iota
	Content
)

type crawInfo struct {
	bookName string
	title    string
	content  string
	next     string
}

var (
	bookName string
	url      string
)

func init() {
	flag.StringVar(&bookName, "bookName", "", "指定书名")
	flag.StringVar(&url, "url", "", "指定抓取的起始链接")
}

func main() {
	flag.Parse()
	if bookName == "" || url == "" {
		flag.Usage()
		return
	}
	result := make(chan string)
	f, err := os.Create(fmt.Sprintf("%s.txt", bookName))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	go LoopCraw(url, bookName, result)
	for context := range result {
		_, _ = f.WriteString(context)
	}
	fmt.Printf("%s 抓取完成\n", bookName)
}

func LoopCraw(url, bookName string, result chan string) {
	for {
		crawInfo, err := craw(url)
		if err != nil {
			log.Fatal(err)
		}
		if crawInfo.bookName != bookName {
			close(result)
			break
		}
		fmt.Printf("抓取<%s>, %s...\n", strings.TrimSpace(crawInfo.title), url)

		result <- crawInfo.title
		result <- crawInfo.content
		url = crawInfo.next
	}
}

func craw(url string) (*crawInfo, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "cf_clearance",
		Value: "egq2KKOXVLJIC8HGgF5yaa7242v4fE5lJiZXJurdYIo-1709703795-1.0.1.1-wnJXYYNm71U2SUXa_IbkrMRwlC1kHLztpUo_D9147Zn7Cu5C51T9LLlQDXEkJPhdWQ0L4eMhVD20cqt4U8GLKA",
	})
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	doc, err := htmlquery.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}

	crawInfo := &crawInfo{}

	// bookName
	node := htmlquery.Find(doc, "//span[@class='post-labels']/a/text()")
	text := htmlquery.InnerText(node[0])
	crawInfo.bookName = strings.TrimSpace(text)

	// title
	node = htmlquery.Find(doc, "//h3[@class='post-title entry-title']/text()")
	text = htmlquery.InnerText(node[0])
	crawInfo.title = fmt.Sprintf("\n\n%s\n\n", strings.TrimSpace(text))

	// content
	node = htmlquery.Find(doc, "//div[@class='post-body entry-content']")
	text = htmlquery.InnerText(node[0])
	builder := strings.Builder{}
	for _, v := range strings.Split(text, "　") {
		builder.WriteString(fmt.Sprintf("　　%s\n", strings.TrimSpace(v)))
	}
	crawInfo.content = builder.String()

	// next
	node = htmlquery.Find(doc, "//a[@class='blog-pager-older-link']/@href")
	text = htmlquery.InnerText(node[0])
	crawInfo.next = strings.TrimSpace(text)

	return crawInfo, nil
}
