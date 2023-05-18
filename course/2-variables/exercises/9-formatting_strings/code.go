package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	var msg string = fmt.Sprintf("Hi %v, your open rate is %.1f percent", name, openRate) // ?

	// don't edit below this line

	fmt.Println(msg)
}
