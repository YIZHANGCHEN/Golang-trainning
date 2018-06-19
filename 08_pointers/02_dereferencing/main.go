package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a) //43
	fmt.Println(&a) //地址

	var b *int = &a
	fmt.Println(b) //地址
	fmt.Println(*b) //43
}

//要显示该地址的值，在地址前+ *
