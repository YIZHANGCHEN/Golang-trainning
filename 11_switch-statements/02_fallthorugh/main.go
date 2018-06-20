package main

import "fmt"

func main() {
	switch "abc" {
	case "gg":
		fmt.Println("Whatsup gg")
	case "MM":
		fmt.Println("Whatsup MM")
	case "abc":
		fmt.Println("Whatsup abc")
		fallthrough //跳过判断，直接打印下一个
	case "1":
		fmt.Println("Whatsup 1")
		fallthrough
	case "2":
		fmt.Println("Whatsup 2")
	case "gvc":
		fmt.Println("Whatsup gvc")
	}
}