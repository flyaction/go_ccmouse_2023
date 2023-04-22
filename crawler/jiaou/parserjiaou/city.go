package parserjiaou

import (
	"regexp"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

const cityRe = `<a class="name" href="(http://www.7799520.com/user/[0-9]+.html)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {

		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})

	}

	return result

}
