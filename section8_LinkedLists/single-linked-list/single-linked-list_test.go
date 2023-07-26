package single_linked_list

import (
	"testing"
)

func TestAppend(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	if l.head.data != "hello" {
		t.Errorf("Expected head to be 'hello', got '%s'", l.head.data)
	}
	if l.tail.data != "hello" {
		t.Errorf("Expected tail to be 'hello', got '%s'", l.tail.data)
	}
	if l.length != 1 {
		t.Errorf("Expected length to be 1, got %d", l.length)
	}

	l.Append("world")
	if l.head.data != "hello" {
		t.Errorf("Expected head to be 'hello', got '%s'", l.head.data)
	}
	if l.tail.data != "world" {
		t.Errorf("Expected tail to be 'world', got '%s'", l.tail.data)
	}
	if l.length != 2 {
		t.Errorf("Expected length to be 2, got %d", l.length)
	}
}

func TestPrepend(t *testing.T) {
	l := NewSingleLinkedList()
	l.Prepend("hello")
	if l.head.data != "hello" {
		t.Errorf("Expected head to be 'hello', got '%s'", l.head.data)
	}
	if l.tail.data != "hello" {
		t.Errorf("Expected tail to be 'hello', got '%s'", l.tail.data)
	}
	if l.length != 1 {
		t.Errorf("Expected length to be 1, got %d", l.length)
	}

	l.Prepend("world")
	if l.head.data != "world" {
		t.Errorf("Expected head to be 'world', got '%s'", l.head.data)
	}
	if l.tail.data != "hello" {
		t.Errorf("Expected tail to be 'hello', got '%s'", l.tail.data)
	}
	if l.length != 2 {
		t.Errorf("Expected length to be 2, got %d", l.length)
	}

}

func TestInsert(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Insert("world", 1)
	if l.head.data != "hello" {
		t.Errorf("Expected head to be 'hello', got '%s'", l.head.data)
	}
	if l.tail.data != "world" {
		t.Errorf("Expected tail to be 'world', got '%s'", l.tail.data)
	}
	if l.length != 2 {
		t.Errorf("Expected length to be 2, got %d", l.length)
	}
}
func TestInsertNode_EmptyList(t *testing.T) {
	l := NewSingleLinkedList()
	l.Insert("hello", 0)
	if l.head.data != "hello" {
		t.Errorf("Expected head to be 'hello', got '%s'", l.head.data)
	}
	if l.tail.data != "hello" {
		t.Errorf("Expected tail to be 'hello', got '%s'", l.tail.data)
	}
	if l.length != 1 {
		t.Errorf("Expected length to be 1, got %d", l.length)
	}
}

func TestInsertNode_AtHead(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Insert("world", 0)
	if l.head.data != "world" {
		t.Errorf("Expected head to be 'world', got '%s'", l.head.data)
	}
	if l.tail.data != "hello" {
		t.Errorf("Expected tail to be 'hello', got '%s'", l.tail.data)
	}
	if l.length != 2 {
		t.Errorf("Expected length to be 2, got %d", l.length)
	}
}

func TestInsertNode_AtTail(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Insert("world", 1)
	if l.head.data != "hello" {
		t.Errorf("Expected head to be 'hello', got '%s'", l.head.data)
	}
	if l.tail.data != "world" {
		t.Errorf("Expected tail to be 'world', got '%s'", l.tail.data)
	}
	if l.length != 2 {
		t.Errorf("Expected length to be 2, got %d", l.length)
	}
}

func TestInsertNode_OutOfBounds(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	err := l.Insert("world", 2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestDelete(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Append("world")
	l.Delete(1)
	if l.head.data != "hello" {
		t.Errorf("Expected head to be 'hello', got '%s'", l.head.data)
	}
	if l.tail.data != "hello" {
		t.Errorf("Expected tail to be 'hello', got '%s'", l.tail.data)
	}
	if l.length != 1 {
		t.Errorf("Expected length to be 1, got %d", l.length)
	}
}
func TestDelete_EmptyList(t *testing.T) {
	l := NewSingleLinkedList()
	err := l.Delete(0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestDelete_Head(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	err := l.Delete(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if l.head != nil {
		t.Errorf("Expected head to be nil, got '%s'", l.head.data)
	}
}

func TestDelete_Tail(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Append("world")
	err := l.Delete(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if l.tail != l.head {
		t.Errorf("Expected tail to be head, got '%s'", l.tail.data)
	}
}

func TestDelete_OutOfBounds(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	err := l.Delete(2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestTraverseToIndex_EmptyList(t *testing.T) {
	l := NewSingleLinkedList()
	_, err := l.traverseToIndex(0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestTraverseToIndex_Head(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	node, err := l.traverseToIndex(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if node.data != "hello" {
		t.Errorf("Expected data to be 'hello', got '%s'", node.data)
	}
}

func TestTraverseToIndex_Tail(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Append("world")
	node, err := l.traverseToIndex(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if node.data != "world" {
		t.Errorf("Expected data to be 'world', got '%s'", node.data)
	}
}

func TestTraverseToIndex_OutOfBounds(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	_, err := l.traverseToIndex(2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPrintList(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Append("world")
	vals := l.PrintList()
	if len(vals) != 2 {
		t.Errorf("Expected 2 values, got %d", len(vals))
	}
	if vals[0] != "hello" {
		t.Errorf("Expected first value to be 'hello', got '%s'", vals[0])
	}
	if vals[1] != "world" {
		t.Errorf("Expected second value to be 'world', got '%s'", vals[1])
	}
}

func TestSearch(t *testing.T) {
	l := NewSingleLinkedList()
	l.Append("hello")
	l.Append("world")
	index := l.Search("hello")
	if index != 0 {
		t.Errorf("Expected index to be 0, got %d", index)
	}
	index = l.Search("world")
	if index != 1 {
		t.Errorf("Expected index to be 1, got %d", index)
	}
}
