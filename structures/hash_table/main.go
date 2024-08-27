package main

import "fmt"

const ArraySize = 5

type bucketNode struct {
	key  string
	next *bucketNode
}

type bucket struct {
	head   *bucketNode
	length int
}

type HashTable struct {
	array [ArraySize]*bucket
}

func InitHashTable() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}

	return hashTable
}

func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.array[index].insert(key)
}

func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.array[index].search(key)
}

func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.array[index].delete(key)
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		node := &bucketNode{key: key}
		node.next = b.head
		b.head = node
		b.length++
	}
}

func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	node := b.head
	for node.next != nil {
		if node.next.key == key {
			node.next = node.next.next
		}
		node = node.next
	}
}

func hash(key string) int {
	var sum int
	for _, character := range key {
		sum += int(character)
	}

	return sum % ArraySize
}

func main() {
	hashTable := InitHashTable()
	hashTable.Insert("ERIC")
	hashTable.Insert("KENNY")
	hashTable.Insert("KYLE")
	hashTable.Insert("STAN")
	hashTable.Insert("RANDY")
	hashTable.Insert("BUTTERS")
	hashTable.Insert("TOKEN")

	fmt.Println(hashTable.Search("ERIC"))
	hashTable.Delete("ERIC")
	fmt.Println(hashTable.Search("ERIC"))
}
