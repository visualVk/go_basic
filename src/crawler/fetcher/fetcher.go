package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// limit request rate
var rateLimiter = time.Tick(time.Microsecond * 10)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bufReader := bufio.NewReader(resp.Body)
	utf8Reader := determineEncoding(bufReader)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(reader *bufio.Reader) io.Reader {
	firstContent, err := bufio.NewReader(reader).Peek(1024)
	var e encoding.Encoding
	if err != nil {
		e = unicode.UTF8
	} else {
		e, _, _ = charset.DetermineEncoding(firstContent, "")
	}
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	return utf8Reader
}
