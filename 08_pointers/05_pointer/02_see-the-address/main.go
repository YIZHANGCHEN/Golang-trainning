package main

import "fmt"

func zero(z *int) {
	fmt.Println(z) //地址
	fmt.Println(*z) //5
	*z = 0
	fmt.Println(z) //地址
	fmt.Println(*z) //0
}

func main() {
	x := 5
	fmt.Println(&x) //地址
	zero(&x)
	fmt.Println(x) //0
}
