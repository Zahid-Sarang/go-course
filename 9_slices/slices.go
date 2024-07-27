package main

import "fmt"

// slice (dynamic array)
// most used construct in go
// + useful methods

func main() {

	// uninitialized slice in nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)

	// capacity => maximum numbers of elements can fit

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)
	// nums = append(nums, 10)
	// nums = append(nums, 11)

	// fmt.Println(cap(nums))
	// fmt.Println(nums)
	// fmt.Println((len(nums)))

	// copy function
	// var slice1 = make([]int, 0, 5)
	// slice1 = append(slice1, 2)
	// var slice2 = make([]int, len(slice1))
	// copy(slice2, slice1)
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// slice operators
	// var num = []int{1, 2, 3}
	// fmt.Println(num[0:2])
	// fmt.Println(num[0:])

	// slice package

	// var num1 = []int{1, 2, 2}
	// var num2 = []int{1, 2}
	// fmt.Println(slices.Equal(num1, num2))

	var TowdSlices = [][]int{{1, 2, 3}, {1, 2, 3}}

	fmt.Println(TowdSlices)
}
