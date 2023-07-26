package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Big O")

	findNemo([]string{"nemo"})                                       // O(n) --> Linear Time   // The number of inputs is the same as the number of operations.
	compressFirstBox([]string{"nemo", "dory"})                       // O(1) --> Constant Time // No matter the number of inputs the number of operations is the same.
	compressFirstTwoBox([]string{"nemo", "dory", "bruce", "marlin"}) // O(2) --> Constant Time // But we round this down to 0(1) because it's a constant.
}

func findNemo(array []string) {
	// O(n) --> Linear Time
	for i := 0; i < len(array); i++ {
		if array[i] == "nemo" {
			fmt.Println("Found Nemo")
		}
	}
}

func compressFirstBox(box []string) {
	// 0(1) --> Constant Time
	fmt.Println(box[0])
}
func compressFirstTwoBox(box []string) {
	// 0(2) --> Constant Time
	// ... but we round this down to 0(1) because it's a constant
	fmt.Println(box[0]) // 0(1)
	fmt.Println(box[2]) // 0(1)
}

// think about how many times each line of code is going to run. this is how you can see the Big O notion
// Technically the Big O is (4 + 3n) but we round this down to 0(n) because the n is the most important part of the equation.
func funChallenge(input []string) int {
	a := 10    // O(1)
	a = 50 + 3 // O(1)

	for i := 0; i < len(input); i++ { // O(n)
		anotherFunction() // O(n) // b/c it's inside the loop
		stranger := true  // O(n) // b/c it's inside the loop
		a++               // O(n) // b/c it's inside the loop
		_ = stranger
	}
	return a // O(1)
}
func anotherFunction() {}

// O(4 + 5n) --> O(5n) --> O(n)
// this would be O(5n) but simplified it turns into O(n) b/c it is still linear time
func funChallenge2(input []string) {
	a := 10    // O(1)
	a = 50 + 3 // O(1)
	a = 3 + 43 // O(1)

	for i := 0; i < len(input); i++ { // O(n)
		b := 1    // O(n) // b/c it's inside the loop
		c := true // O(n) // b/c it's inside the loop
		a++       // O(n) // b/c it's inside the loop
		_ = a
		_ = b
		_ = c
	}
	for i := 0; i < len(input); i++ { // O(n)
		b := 1    // O(n) // b/c it's inside the loop
		c := true // O(n) // b/c it's inside the loop
		_ = a
		_ = b
		_ = c
	}
	_ = "i don't know" // O(1)
}

// it is not 0(n) b/c we have 2 inputs. it is 0(a + b)
// # 3rd Big O Rule Different terms for inputs
func compressBoxesTwice(boxes []string, boxes2 []string) {
	// O(a + b)
	for i := 0; i < len(boxes); i++ {
		fmt.Println(boxes[i])
	}
	for i := 0; i < len(boxes2); i++ {
		fmt.Println(boxes2[i])

	}
}

func LogAllPairs(array []string) {
	// O(n^2) --> Quadratic Time
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Println(array[i], array[j])
		}
	}
}
