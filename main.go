package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"crawler/lib"
)



// Extract all http** links from a given webpage
func crawl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns
	log.Println("=====================================")
	bodyByte, _ := ioutil.ReadAll(b)
	resStr := string(bodyByte)
	//log.Println(resStr)
	lib.WriteFile(resStr,url)

}

func main() {

	//1. Extract url and download html to local.
	seedUrls := []string {"https://docs.docker.com/","http://daily.zhihu.com/"}

	for _, url := range seedUrls {
		crawl(url)
	}

	//2. parse html sample.
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	lib.Parse(s)

}