package main

import "fmt"

const ArraySize = 5

type node struct {
	key  string
	next *node
}

type bucket struct {
	head   *node
	length int
}

type HashTable struct {
	buckets [ArraySize]*bucket
}

func InitHashTable() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.buckets {
		hashTable.buckets[i] = &bucket{}
	}

	return hashTable
}

func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.buckets[index].insert(key)
}

func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.buckets[index].search(key)
}

func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.buckets[index].delete(key)
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		node := &node{key: key}
		node.next = b.head
		b.head = node
		b.length++
	}
}

func (b *bucket) search(key string) bool {
	node := b.head
	for node != nil {
		if node.key == key {
			return true
		}
		node = node.next
	}
	return false
}

func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		b.length--
		return
	}

	node := b.head
	for node.next != nil {
		if node.next.key == key {
			node.next = node.next.next
			b.length--
			return
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

	fmt.Println(hashTable.Search("TOKEN"))
	hashTable.Delete("TOKEN")
	fmt.Println(hashTable.Search("TOKEN"))
}
