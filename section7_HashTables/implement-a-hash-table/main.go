package main

import "fmt"

// implement_a_hash_table

type HashSet struct {
	//hash      int // this is redundant b/c the index is the hash value
	hashPairs []HashPair
}

type HashPair struct {
	key   string
	value int
}

type HashTable struct {
	//hashSets []any // we need to be able to store both the key and value
	//hashSets [][]int // we need to allow for the retreival of keys even if there was a hash collision
	hashSets []HashSet
}

func NewHashTable(size int) *HashTable {

	return &HashTable{
		//hashSets: make([]any, 0, size),
		//hashSets: make([][]int, size, size),
		hashSets: make([]HashSet, size, size),
	}
}

// this function will return a index memory space based on the key in the hash algorithm
func (h *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % cap(h.hashSets)
	}
	return hash
} // O(1)

func (h *HashTable) Set(key string, value int) {
	hash := h.hash(key)

	hashPair := HashPair{
		key:   key,
		value: value,
	}

	// if the hash key is new, initialize value with fresh array
	if h.hashSets[hash].hashPairs == nil {
		h.hashSets[hash].hashPairs = make([]HashPair, 0, 1)
	}

	h.hashSets[hash].hashPairs = append(h.hashSets[hash].hashPairs, hashPair)
} // O(1)

func (h *HashTable) Get(key string) int {
	hash := h.hash(key)

	if len(h.hashSets[hash].hashPairs) > 1 {
		for _, pair := range h.hashSets[hash].hashPairs {
			if pair.key == key {
				return pair.value
			}
		}
	} else {
		return h.hashSets[hash].hashPairs[0].value
	}

	return -1
} // O(1)... but O(n) if there was a collision

// return all the keys
func (h *HashTable) keys() []string {
	keys := []string{}

	for _, set := range h.hashSets {
		if set.hashPairs != nil {
			for _, pair := range set.hashPairs {
				keys = append(keys, pair.key)
			}
		}
	}

	return keys
}

// return all the values
func (h *HashTable) values() []int {
	values := []int{}

	for _, set := range h.hashSets {
		if set.hashPairs != nil {
			for _, pair := range set.hashPairs {
				values = append(values, pair.value)
			}
		}
	}

	return values
}

func main() {
	myHashTable := NewHashTable(50)
	myHashTable.Set("grapes", 10000)
	myHashTable.Set("grapes", 20000)
	myHashTable.Set("apples", 500)
	fmt.Println(myHashTable.Get("grapes"))
	fmt.Println(myHashTable.hashSets)
	fmt.Println(myHashTable.keys())
	fmt.Println(myHashTable.values())
}
