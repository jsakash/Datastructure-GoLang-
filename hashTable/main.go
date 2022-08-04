package main

import (
	"fmt"
)

// Array size
const Size = 7

// HashTable will hold an array
type HashTable struct {
	array [Size]*bucket
}

// bucket is a linked list in each slot of the array
type bucket struct {
	head *bucketNode
}

// bucketNode Structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// insert will create a bucket in each slot of the hash table
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

// delete will take a key and delete from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take a key,create a node with the key and insert  the mode in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Aleady exists")
	}

}

// search will take in a key and return true if the bucket has the key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// Hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % Size
}

// init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {

	hashTable := Init()
	list := []string{
		"ASHNA",
		"AKASH",
		"ABHISHEK",
		"NIKHIL",
		"ALAN",
		"AROMAL",
		"ASWIN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("ASWIN")
	fmt.Println("ALAN", hashTable.Search("ALAN"))
	fmt.Println("ASWIN", hashTable.Search("ASWIN"))

}
