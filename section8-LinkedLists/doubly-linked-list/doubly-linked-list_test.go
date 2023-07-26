package main

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")
	list.Append("!!!")

	expected := []string{"hi", "there", "!!!"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPrepend(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Prepend("hi")
	list.Prepend("there")
	list.Prepend("!!!")

	expected := []string{"!!!", "there", "hi"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsert(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")

	list.Insert("world", 1)

	expected := []string{"hi", "world", "there"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDelete(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")
	list.Append("!!!")

	list.Delete(1)

	expected := []string{"hi", "!!!"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSearch(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")
	list.Append("!!!")

	index := list.Search("there")

	if index != 1 {
		t.Errorf("Expected index 1, got %v", index)
	}
}

func TestTraverseToIndex(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")
	list.Append("!!!")

	node, err := list.traverseToIndex(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if node.data != "there" {
		t.Errorf("Expected data 'there', got %v", node.data)
	}
}

func TestLength(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")
	list.Append("!!!")

	length := list.Length()

	if length != 3 {
		t.Errorf("Expected length 3, got %v", length)
	}
}

func TestAppendEmptyList(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")

	expected := []string{"hi"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPrependEmptyList(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Prepend("hi")

	expected := []string{"hi"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertAtHead(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")

	list.Insert("world", 0)

	expected := []string{"world", "hi", "there"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertAtTail(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")

	list.Insert("world", 2)

	expected := []string{"hi", "there", "world"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeleteHead(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")

	list.Delete(0)

	expected := []string{"there"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeleteTail(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")

	list.Delete(1)

	expected := []string{"hi"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertAndDelete(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append("hi")
	list.Append("there")
	list.Append("!!!")

	list.Insert("world", 0)
	list.Delete(1)
	list.Insert("hello", 2)

	expected := []string{"world", "hello", "!!!"}
	actual := list.PrintList()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Verify that the next and previous pointers are pointing to the correct nodes.
	node := list.head
	for i := 0; i < len(expected); i++ {
		if i == 0 {
			if node.next != list.head.next {
				t.Errorf("Expected next pointer to point to head, got %v", node.next)
			}
		} else {
			n, err := list.traverseToIndex(i)
			if err != nil {
				t.Errorf("ERROR")
			}
			if node.next != n {
				t.Errorf("Expected next pointer to point to node %v, got %v", n, node.next)
			}

			n, err = list.traverseToIndex(i - 1)
			if err != nil {
				t.Errorf("ERROR")
			}
			if node.previous != n {
				t.Errorf("Expected previous pointer to point to node %v, got %v", n, node.previous)
			}
		}
		node = node.next
	}
}
