package main

func containsDuplicate(nums []int) bool {
	//return containsDuplicate_BruteForce(nums)
	return containsDuplicate_Better(nums)
}

func containsDuplicate_Better(nums []int) bool {

	records := make(map[int]interface{}) // we don't car about the value, only the key
	for _, num := range nums {
		if _, exists := records[num]; exists {
			return true
		} else {
			records[num] = nil
		}
	}

	return false
}
func containsDuplicate_BruteForce(nums []int) bool {

	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				return true
			}
		}
	}
	return false
}
