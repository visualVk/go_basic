package parser

import (
	"example.com/ch4/src/crawler/engine"
	"example.com/ch4/src/crawler/model"
	"fmt"
	"regexp"
	"strconv"
)
var nameRe = regexp.MustCompile(`<th><a[^>]*>([^<]*)</a></th>`)
var ageRe = regexp.MustCompile(`<td[^>]*><span[^>]*>年龄：</span>(\d+)</td>`)
var marriageRe = regexp.MustCompile(`<td[^>]*><span[^>]*>婚况：</span>([^<]*)</td>`)

func ParseProfile(contents []byte) engine.ParserResult{
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe)
	profile.Name = extractString(contents, nameRe)

	result := engine.ParserResult{}
	result.Items = append(result.Items, profile)

	fmt.Printf("user: %s, age: %d, Marriage: %s\n", profile.Name, profile.Age, profile.Marriage)

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string{
	matches := re.FindSubmatch(contents)

	if len(matches) >= 2{
		return string(matches[1])
	}else {
		return ""
	}

}