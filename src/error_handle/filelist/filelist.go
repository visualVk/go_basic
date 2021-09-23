package filelist

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const CURWD = "src/error_handle/"

const prefix = "/list/"

type userErr string

func (e userErr) Error() string {
	return e.Message()
}

func (e userErr) Message() string {
	return string(e)
}

func Handler(w http.ResponseWriter, r *http.Request) error {
	// return userErr when url dosn't contain "/list/"
	index := strings.Index(r.URL.Path, prefix)
	if index == -1 {
		return userErr("path must start with" + prefix)
	}

	path := r.URL.Path[len("/list/"):]
	path = CURWD + path
	fmt.Println("file name: ", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	w.Write(all)
	return nil
}
