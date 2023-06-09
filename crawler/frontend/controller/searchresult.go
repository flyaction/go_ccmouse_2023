package controller

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"regexp"

	"imooc.com/ccmouse/learngo/crawler/engine"

	"imooc.com/ccmouse/learngo/crawler/frontend/model"

	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"
	"imooc.com/ccmouse/learngo/crawler/frontend/view"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandle(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}

}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s,from=%d", q, from)
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

const pageSize = 10

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {

	var result model.SearchResult

	result.Query = q

	resp, err := h.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from

	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	if result.Start == 0 {
		result.PrevFrom = -1
	} else {
		result.PrevFrom = (result.Start - 1) / pageSize * pageSize
	}
	result.NextFrom = result.Start + len(result.Items)

	log.Printf("+%v", result.Items)

	return result, nil

}

// Rewrites query string. Replaces field names
// like "Age" to "Payload.Age"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
