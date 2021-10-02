package main

import (
	"log"
	"net/http"
	"os"

	"example.com/ch4/src/error_handle/filelist"
)

// how to handle error/panic
// 1. define function for dealing with error/panic which return value is error(can be error customed)
// 	- error customed should be implementing interface of Error
// 2. define function for dealing with above function's return.
// 	- use type assertion to tell error

// custom error which have inteface of error and function named Message
type userError interface {
	error
	Message() string
}

// alias func with appHandler for defining parameter
type appHandler func(http.ResponseWriter, *http.Request) error

// handle filelist.Hanlder with dealing with error thrown
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := filelist.Handler(w, r)
		if err != nil {
			log.Printf("Error handle request %s", err.Error())

			if userErr, ok := err.(userError); ok { // type assertion
				http.Error(w, userErr.Message(), http.StatusInternalServerError)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(w, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelist.Handler))
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
