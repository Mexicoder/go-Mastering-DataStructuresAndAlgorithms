package google_video_interview_question

import "fmt"

// lets assume the data is sorted
// we have small numbers on 1 side and large numbers on the other
// Time Complexity: O(n)
//
//	b/c worst case it has to traverse all elements in data
func HasPairWithSum_sortedData(data []int, sum int) bool {
	low, high := 0, len(data)-1 // need to remember that there is a zero index

	for low < high {

		if data[low]+data[high] == sum {
			// Success
			fmt.Println("found it")
			return true
		} else if data[low]+data[high] < sum {
			// too small
			fmt.Println("too small")
			low++
		} else { // } else if data[low]+data[high] > sum {
			// too high
			fmt.Println("too high")
			high--
		}
	}
	return false
}

// lets assume the data is not sorted
func HasPairWithSum_notSortedData(data []int, sum int) bool {

	//past := make(map[int]struct{}, 0)
	past := make(map[int]struct{}, len(data)) // ??? can i do this for constant memory complexity?

	for _, datum := range data {
		need := sum - datum
		if _, exists := past[need]; exists {
			// found what we need
			fmt.Println("datum: ", datum, ", past val: ", need)
			return true
		} else {
			if datum > sum {
				// if the number is already too big i don't want to keep track of it
				continue
			} else {
				// didn't find it so adding it to past
				past[datum] = struct{}{} // i don't think the value means anything here?
			}
		}
	}

	return false
}

func containsCommonItem1(arr1 []string, arr2 []string) bool {
	for _, s := range arr1 {
		for _, s2 := range arr2 {
			if s == s2 {
				return true
			}
		}
	}
	return false
}

func containsCommonItem2(arr1 []string, arr2 []string) bool {
	map1 := make(map[string]struct{}, len(arr1))
	for _, s := range arr1 {
		map1[s] = struct{}{}
	}

	for _, s := range arr2 {
		if _, exists := map1[s]; exists {
			return true
		}
	}
	return false
}
