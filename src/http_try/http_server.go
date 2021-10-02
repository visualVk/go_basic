package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respContent, err2 := httputil.DumpResponse(resp, true)
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("%s\n", respContent)
}
