package parserjiaou

import (
	"regexp"
	"strconv"

	"imooc.com/ccmouse/learngo/crawler/config"

	"imooc.com/ccmouse/learngo/crawler/model"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

var ageRe = regexp.MustCompile(
	`<span class="age s1">年龄：</span>(\d+)岁</span>`)
var heightRe = regexp.MustCompile(
	`<li>身高：<span>(\d+)cm</span></li>`)
var incomeRe = regexp.MustCompile(
	`<li>收入：<span>([^<]+)</span></li>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<li>星座：<span>([^<]+)</span></li>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<span class="education">([^<]+)</span></p>`)
var occupationRe = regexp.MustCompile(
	`<li>职业：<span>([^<]+)</span></li>`)
var hokouRe = regexp.MustCompile(
	`<li>籍贯：<span>([^<]+)</span></li>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`www.7799520.com/user/([\d]+).html`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {

	profile := model.Profile{}

	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "jiaou",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(
	contents []byte,
	url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (
	name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}

}

//func ProfileParser(name string) engine.ParserFunc {
//
//	return func(c []byte, url string) engine.ParseResult {
//		return ParseProfile(c, url, name)
//	}
//
//}
