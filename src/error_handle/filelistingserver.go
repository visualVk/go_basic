package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const CURWD = "src/error_handle/"

func main() {
	http.HandleFunc("/list/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/list/"):]
		path = CURWD + path
		fmt.Println("file name: ", path)
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		w.Write(all)
	})

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
