package main

import "fmt"

func add(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	var sum int = 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

func configs(nums ...int) {
	if len(nums) > 0 {
		fmt.Println(nums[0])
	} else {
		fmt.Println("Default")
	}

}

func main() {
	fmt.Println(add())
	fmt.Println(add(1))
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))

	configs()
	configs(1, 2, 3)
}
