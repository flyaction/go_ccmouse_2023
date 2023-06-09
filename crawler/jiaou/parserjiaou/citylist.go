package parserjiaou

import (
	"regexp"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

const cityListRe = `<a href="(http://www.7799520.com/jiaou/[0-9a-z]+)">([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	//limit := 3

	for _, m := range matches {

		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})

		//limit--
		//if limit <= 0 {
		//	break
		//}

	}

	return result

}
