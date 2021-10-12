package parser_test

import (
	"example.com/ch4/src/crawler/engine"
	parser "example.com/ch4/src/crawler/parser/citylist"
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	file, err := ioutil.ReadFile("./zhenghun.out")
	if err != nil {
		return
	}

	//fmt.Printf("%s\n", file)

	result := parser.ParserCityList(file)
	table := []struct {
		request engine.ParserResult
		len     int
	}{
		{
			request: result,
			len:     470,
		},
	}

	for _, data := range table {
		requestLen := len(data.request.Requests)
		itemLen := len(data.request.Items)
		if itemLen != data.len || requestLen != data.len {
			t.Errorf("expect len is %d, but requestLen is %d, itemLen is %d", data.len, requestLen, itemLen)
		}
	}
}
