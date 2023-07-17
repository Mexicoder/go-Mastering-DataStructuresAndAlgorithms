package main

import (
	"errors"
	"fmt"
)

type MyArray struct {
	length int
	data   []interface{}
}

func NewMyArray() *MyArray {
	return &MyArray{
		length: 0,
		data:   make([]interface{}, 0),
	}
}

func (a *MyArray) Get(index int) (interface{}, error) {
	// 0 0
	if index >= a.length {
		return nil, errors.New("out of bounds")
	}
	return a.data[index], nil
}

func (a *MyArray) Push(item interface{}) {
	a.data = append(a.data, item)
	a.length += 1
	fmt.Println("Pushed: ", item)
	fmt.Println("Length now: ", a.length)
}

func (a *MyArray) Pop() interface{} {
	val := a.data[len(a.data)-1]

	a.data = a.data[:len(a.data)-1]
	a.length -= 1
	fmt.Println("Popped: ", val)

	return val
}

func (a *MyArray) Delete(index int) {

}

// -----------------------------------

func main() {
	newArray := NewMyArray()

	fmt.Println(newArray.Get(0))

	newArray.Push("hi")

	fmt.Println(newArray.Get(0))

	newArray.Push("you")

	newArray.Push("!")

	fmt.Println(newArray)

	fmt.Println(newArray.Pop())

	fmt.Println(newArray)

}
