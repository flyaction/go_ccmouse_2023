package parserjiaou

import (
	"regexp"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

const cityRe = `<a class="name" href="(http://www.7799520.com/user/[0-9]+.html)" target="_blank">([^<]+)</a>`

var (
	profileRe = regexp.MustCompile(`<a class="name" href="(http://www.7799520.com/user/[0-9]+.html)" target="_blank">([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`<a target="_blank" href="(http://www.7799520.com/jiaou/shandong/[a-z]+)">([^<]+)</a>`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {

		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})

	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result

}
