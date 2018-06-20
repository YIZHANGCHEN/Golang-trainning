package main

import "fmt"

func main() {
	switch "GG" {
	case "abc":
		fmt.Println("Whatsup abc")
	case "MM":
		fmt.Println("Whatsup MM")
	case "GG":
		fmt.Println("Whatsup GG")
	default:
		fmt.Println("where is your friends?")
	}
}