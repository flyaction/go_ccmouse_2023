package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID       string  `json:"id"`
	Name     string  `json:"name,omitempty"` //omitempty为空不展示
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
}

func main() {
	//marshal()
	//unMarshal()
	parseNLP()
}

func marshal() {
	o := Order{
		ID:         "1234",
		TotalPrice: 30,
		Items: []OrderItem{
			{ID: "item_1", Name: "learn go", Price: 15, Quantity: 2},
			{ID: "item_2", Name: "learn go", Price: 1, Quantity: 1},
		},
	}

	//fmt.Printf("%+v\n", o)

	b, err := json.Marshal(o)
	if err != nil {
		panic(b)
	}
	fmt.Printf("%s\n", b)
}

func unMarshal() {

	s := `{"id":"1234","items":[{"id":"item_1","name":"learn go","price":15,"quantity":2},{"id":"item_2","name":"learn go","price":1,"quantity":1}],"total_price":30}`
	var o Order

	err := json.Unmarshal([]byte(s), &o)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", o)

}

func parseNLP() {
	res := `{
"data":[
  {
    "span": "斗鱼",
    "label": "111",
    "score": 0.9068
  },
 {
    "span": "虎牙",
    "label": "222",
    "score": 0.9068
  },
{
    "span": "熊猫",
    "label": "222",
    "score": 0.9068
  }
]}`
	//m := make(map[string]interface{})

	m := struct {
		Data []struct {
			Span string `json:"span"`
		} `json:"data"`
	}{}

	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m)
	//fmt.Printf("%+v\n", m["data"].([]interface{})[1].(map[string]interface{})["span"])
	fmt.Printf("%+v\n", m.Data[2].Span)
}
