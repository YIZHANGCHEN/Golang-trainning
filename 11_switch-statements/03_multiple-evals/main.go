package main

import "fmt"

func main() {
	switch "TT" {
	case "TT", "JJ":
		fmt.Println("Whatsup TT, or, err, JJ")
	case "EDC", "ETT":
		fmt.Println("Both of your names start with E")
	case "Julian", "Sushant":
		fmt.Println("Whatsup Julian / Sushant")
	}
}