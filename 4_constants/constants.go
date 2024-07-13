package main

import "fmt"

const age = 22
const married bool = true

func main() {

	fmt.Println(age)
	const name = "golang"
	// you can't change variables

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

}
