package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "reverse this string" // 19 chars

	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}

	fmt.Println()
	fmt.Println("len(str): ", len(str))

	// len(str) - 1 // b/c len(str) is 19 1 indexed and we need the 0 index version... so we -1
	// or else is will try and get str[19] which is "index of our range" // IMPORTANT to remember
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Println()

	str2 := "abcdefg"
	//arrayVal := []string{str2}

	fmt.Println(reverseSlice2(str2))

	arrayIntVal := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseSlice2(arrayIntVal))

}

func reverseSlice(arr []interface{}) []interface{} {
	// check input
	if len(arr) < 2 {
		return arr
	}

	revs := make([]interface{}, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		revs = append(revs, arr[i])
	}
	return revs
}
func reverseSlice2(arr interface{}) interface{} {
	// check input
	if reflect.ValueOf(arr).Len() < 2 {
		return arr
	}

	revs := make([]interface{}, 0, reflect.ValueOf(arr).Len())

	for i := reflect.ValueOf(arr).Len() - 1; i >= 0; i-- {
		revs = append(revs, reflect.ValueOf(arr).Index(i).Interface())
	}
	return revs
}
