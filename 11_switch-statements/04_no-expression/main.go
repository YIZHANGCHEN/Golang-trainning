package main

import "fmt"

func main() {

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Whatsup my friend with name of length 2")
	case myFriendsName == "Mar":
		fmt.Println("Whatsup Mar")
	case myFriendsName == "JJ":
		fmt.Println("Whatsup JJ")
	case myFriendsName == "MM", myFriendsName == "GG":
		fmt.Println("Your name is either MM or GG")
	case myFriendsName == "Ju":
		fmt.Println("Whatsup Ju")
	case myFriendsName == "Sushant":
		fmt.Println("Whatsup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}