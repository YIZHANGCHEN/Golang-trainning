package main

import "fmt"

func main() {

	greeting := func() { //func 名字 （参数） returns
		fmt.Println("Hello world!")
	}

	greeting()
}