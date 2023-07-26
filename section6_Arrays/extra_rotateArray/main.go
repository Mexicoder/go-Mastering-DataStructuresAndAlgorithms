package main

import "fmt"

func rotate(nums []int, k int) {

	// i could create another array and start from index of k and when i get to the end use math to roll over the index to 0
	// what if i created a placeholder for the value extracted?

	fmt.Println(nums)
	rotated := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		index := len(nums) - k
		if index > len(nums) {
			index -= len(nums)
		}
		rotated[i] = nums[index]
		fmt.Println("i: ", i, " nums[index]: ", nums[index])
	}

	nums = rotated
	//fmt.Println(nums)
}

// start [0,1,2,3]
// k = 1
// end [3,0,1,2]

// i-k
// if negative math.abs -len(nums)
