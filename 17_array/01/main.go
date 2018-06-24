package main

import "fmt"

func main() {
	var x [88]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[87])

	for i := 65; i <= 152; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[87])
}