package main

// I still have a hard time wrapping my head around the logic for the O(n) solution but i do understand it. it just feels a bit too small to see

// extra_max_sumarray
func maxSubArray(nums []int) int {
	// // Brute Force
	//largestSum := math.MinInt32
	////sumString := ""
	//
	//for i := 0; i < len(nums); i++ {
	//	sum := 0
	//	for j := i; j < len(nums); j++ {
	//		sum += nums[j]
	//		if sum > largestSum {
	//			largestSum = sum
	//		}
	//	}
	//}
	//return largestSum

	runningSum := nums[0]
	largestSum := nums[0]

	// i = 1 b/c our cursors are already at index 0
	for i := 1; i < len(nums); i++ {
		// check which is larger. runningSum+nextValue or the current sum
		runningSum = max(runningSum+nums[i], nums[i])
		if largestSum < runningSum {
			largestSum = runningSum
		}
	}

	return largestSum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
