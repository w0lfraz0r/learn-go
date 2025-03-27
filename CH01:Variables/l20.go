package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %0.1f percent\n", name, openRate)

	// don't edit below this line

	fmt.Print(msg)
}
