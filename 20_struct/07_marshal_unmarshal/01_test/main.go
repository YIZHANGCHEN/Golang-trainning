package main

import (
	"fmt"
	"encoding/json"
)

type Station struct {
	Mac string `json:"mac"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Lrrt int `json:"lrrt"`
	Areas []string `json:"areas"`
	Loc struct {
		Lat     float64 `json:"lat"`
		LevelID int     `json:"levelId"`
		Lng     float64 `json:"lng"`
		Prec    int     `json:"prec"`
	} `json:"loc"`
	Rss  struct {
		Seven072CF21EB13 int `json:"7072CF21EB13"`
		Seven072CF4F6862 int `json:"7072CF4F6862"`
	} `json:"rss"`
}



func main() {
	var s1 Station
	fmt.Println(s1.Mac)
	fmt.Println(s1.Name)
	fmt.Println(s1.Desc)
	fmt.Println(s1.Lrrt)
	fmt.Println(s1.Areas)
	fmt.Println(s1.Loc)
	fmt.Println(s1.Rss)

	bs := []byte(`  {
        "mac": "0060B3036C79",
        "name": "Test device 1",
        "desc": "",
        "loc": {
            "lat": 32.79764,
            "lng": -117.249125,
            "levelId": 0,
            "prec": 15
        },
        "lrrt": 3137626,
		"areas": [
            	"5"
        		],
        "rss": {
            "7072CF4F6862": -92,
            "7072CF21EB13": -88
        }
    }`)
	json.Unmarshal(bs, &s1)

	fmt.Println("---------------")
	fmt.Println(s1.Mac)
	fmt.Println(s1.Name)
	fmt.Println(s1.Desc)
	fmt.Println(s1.Lrrt)
	fmt.Println(s1.Areas)
	fmt.Println(s1.Loc)
	fmt.Println(s1.Rss)
	fmt.Printf("%T \n", s1)
	fmt.Println("---------------")
	fmt.Println(s1)
}