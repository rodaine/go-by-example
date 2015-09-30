package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func mpl(v interface{}) {
	o, _ := json.Marshal(v)
	fmt.Println(string(o))
}

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	mpl(true)
	mpl(1)
	mpl(2.34)
	mpl("gopher")
	mpl([]string{"a", "b", "c"})
	mpl(map[string]int{"a": 1, "b": 2})

	mpl(&Response1{Page: 1, Fruits: []string{"a", "b", "c"}})
	mpl(&Response2{Page: 1, Fruits: []string{"a", "b", "c"}})

	byt := []byte(`{"page":6,"fruits":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["page"].(float64)
	fmt.Println(num)

	strs := dat["fruits"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	res := &Response2{}
	json.Unmarshal(byt, res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
