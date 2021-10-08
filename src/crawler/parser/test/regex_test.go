package test

import (
	"io/ioutil"
	"os"
	"testing"

	"example.com/ch4/src/crawler/citylist"
)

func TestRegex(t *testing.T) {
	f, err := os.Open("output.out")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	tableData := []struct {
		content []byte
		len     int
	}{
		{
			content: content,
			len:     470,
		},
	}

	for _, data := range tableData {
		matches := citylist.PrintCityList(data.content)
		if data.len != len(matches) {
			t.Errorf("Expected %d, got %d", data.len, len(matches))
		}
	}
}
