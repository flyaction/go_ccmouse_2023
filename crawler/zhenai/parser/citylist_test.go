package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, _ := ioutil.ReadFile("citylist_test_data.html")

	result := ParseCityList(contents)

	//fmt.Printf("%s\n", contents)

	const ResultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Requests) != ResultSize {
		t.Errorf("result should have %d requests;but had %d", ResultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d:%s;but "+"was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != ResultSize {
		t.Errorf("result should have %d items;but had %d", ResultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d:%s;but "+"was %s", i, city, result.Items[i].(string))
		}
	}
}
