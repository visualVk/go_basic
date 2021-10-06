package main

import (
	"net/http"

	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	transform.NewReader()
}
