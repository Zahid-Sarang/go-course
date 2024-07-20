package main

import "fmt"

func main() {
	// age := 10

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("person is teenager")
	// } else {
	// 	fmt.Println("person is kid")
	// }

	// var role = "admin"
	// var hasPermission = true

	// if role == "admin" && hasPermission {
	// 	fmt.Println("YES")
	// }

	// we can declare the variable inside the if construct
	if age := 20; age >= 18 {
		fmt.Println("allowed for driving licensed")
	} else if age >= 12 {
		fmt.Println("dont allow for driving licensed")
	}

	// Go does not have ternary , you will have to use normal if else
}
