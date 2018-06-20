package main

import "fmt"

func zero(z *int) {
	*z = 0 //value 0 stored in address z
}

func main() {
	x := 5
	zero(&x) //value 5 address stored in address func zero
	fmt.Println(x) // 0
}