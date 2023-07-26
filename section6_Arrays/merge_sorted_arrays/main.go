package main

func MergeSortedArrays(arr1, arr2 []int) []int {

	// Brute force
	// declare empty result array
	// loop through arr1 and arr2 inside of a nested loop, compared index values and push smaller value onto result and increment it's index... do till finish

	result := make([]int, 0, len(arr1)+len(arr2))
	//ai1, ai2 := 0, 0
	// // This doesn't work well b/c this logic currently goes out of bounds
	//for ai1 <= len(arr1) {
	//	for ai2 <= len(arr2) {
	//		if arr1[ai1] < arr2[ai2] {
	//			result = append(result, arr1[ai1])
	//			ai1 += 1
	//		} else {
	//			result = append(result, arr2[ai2])
	//			ai2 += 1
	//		}
	//	}
	//}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i += 1
			//fmt.Println("i: ", i)
		} else {
			result = append(result, arr2[j])
			j += 1
			//fmt.Println("j: ", j)
		}
	}

	if i != len(arr1) {
		result = append(result, arr1[i:]...)
		//fmt.Println("arr1 remainder: ", arr1[i:])
	} else if j != len(arr2) {
		result = append(result, arr2[j:]...)
		//fmt.Println("arr2 remainder: ", arr2[j:])
	}

	return result
}

func main() {

}
