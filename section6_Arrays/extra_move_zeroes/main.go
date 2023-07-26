package main

func moveZeroes(nums []int) {
	//moveZeroes_BruteForce(nums)
	moveZeroes_Better(nums)
}

// the better solution taught me to think a little differently. instead of counting the zeros we can keep track of where non zero numbers stop. by doing this we can loop through the array, and if element is not zero we move it to the end of the non zero index and increment that index. when we are done shifting we run another loop to start form the non zero index and populate the rest with zeros
func moveZeroes_Better(nums []int) {
	l := 0 // keeping track of non zero elements

	for r := range nums {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}

	for ; l < len(nums); l++ {
		nums[l] = 0
	}

	return
}

func moveZeroes_BruteForce(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// shift all elements
			for j := i; j < len(nums)-1; j++ {
				// take the next value and assign it the current value. shifting element to the left 1
				nums[j] = nums[j+1]
			}
			nums[len(nums)-1] = 0
		}
	}
}
