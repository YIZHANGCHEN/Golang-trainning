package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age float64 `json:"wisdom score"`
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":-20.000000001}`)

	fmt.Println("----------")
	fmt.Println(bs)
	fmt.Println("----------")

	json.Unmarshal(bs, &p1)




	fmt.Println("----------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}