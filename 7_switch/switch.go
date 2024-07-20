package main

import "fmt"

func main() {
	// simple switch

	// i := 4

	// switch i {
	// case 1:
	// 	fmt.Println(1)
	// case 2:
	// 	fmt.Println(2)
	// case 3:
	// 	fmt.Println(3)
	// case 4:
	// 	fmt.Println(4)
	// default:
	// 	fmt.Println("unknown", i)
	// }

	// multiple conditions switch

	// switch time.Now().Weekday() {

	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")

	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
		}
	}

	whoAmI(50.892938)

}
