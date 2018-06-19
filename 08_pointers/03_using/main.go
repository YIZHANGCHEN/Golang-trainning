package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a) //43
	fmt.Println(&a) //地址

	var b *int = &a
	fmt.Println(b) //地址
	fmt.Println(*b) //43

	*b = 42 //b地址的值变成42
	fmt.Println(a) //42
}


