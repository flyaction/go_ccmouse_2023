package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "action",
		"age":  "12",
		"sex":  "boy",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	if name, ok := m["name"]; ok {
		fmt.Println("get name ===" + name)
	} else {
		fmt.Println("key name not exist")
	}

	delete(m, "name")

	if name, ok := m["name"]; ok {
		fmt.Println("get name ===" + name)
	} else {
		fmt.Println("key name not exist")
	}

	fmt.Println(len(m))
}
