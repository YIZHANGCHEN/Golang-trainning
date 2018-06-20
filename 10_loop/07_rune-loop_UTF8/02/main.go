package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i))) //如果"i"，则打印为105-i对应的数字，非递增
	}
	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
}
