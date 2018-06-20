package main

import "fmt"

func main() {

	b := true

	if food := "choco"; b { //先定义变量再判定可以
		fmt.Println(food)
	}
}