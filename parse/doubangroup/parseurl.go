package doubangroup

import (
	"github.com/1152545264/goSpyder/collect"
	"regexp"
)

const urlListRe = `(https://www.douban.com/group/topic/[0-9a-z]+/)"[^>]*>([^<]+)</a>`

func ParseURL(contents []byte, req *collect.Request) collect.ParseResult {
	re := regexp.MustCompile(urlListRe)

	matches := re.FindAllSubmatch(contents, -1)
	result := collect.ParseResult{}

	for _, m := range matches {
		u := string(m[1])
		result.Requests = append(result.Requests, &collect.Request{
			Method: "GET",
			Task:   req.Task,
			Url:    u,
			Depth:  req.Depth + 1,
			ParseFunc: func(c []byte, request *collect.Request) collect.ParseResult {
				return GetContent(c, u)
			},
		})
	}
	return result
}

const ContentRes = `<div class="topic-content">[\s\S]*?阳台<div`

func GetContent(contents []byte, url string) collect.ParseResult {
	re := regexp.MustCompile(ContentRes)
	ok := re.Match(contents)
	if !ok {
		return collect.ParseResult{
			Items: []interface{}{},
		}
	}
	result := collect.ParseResult{
		Items: []interface{}{url},
	}
	return result
}
