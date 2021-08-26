package main

import "fmt"

//Arraysize is the size of the HashTable array
const ArraySize = 7

//Hash table structure will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

//Bucket structure for linkedList
type bucket struct {
	head *bucketNode
}

//Bucket Node structure
type bucketNode struct {
	Key  string
	next *bucketNode
}

//Insert will take in a Key and add to the Hashtable array
func (h *HashTable) Insert(s string) {
	index := Hash(s)
	h.array[index].BucketInsert(s)

}

//Search will take in a Key and Return true if the word is present in the hash table
func (h *HashTable) Search(s string) bool {
	index := Hash(s)
	return h.array[index].BucketSearch(s)
}

//Delete will take in a key and delete if it is present in the hashtable
func (h *HashTable) Delete(s string) {
	index := Hash(s)
	h.array[index].BucketDelete(s)
}

//BucketInsert will create a node with the key and insert it in the bucket
func (b *bucket) BucketInsert(k string) {
	if !b.BucketSearch(k) {
		newNode := &bucketNode{Key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("%s is already exist\n", k)
	}

}

//BucketSearch will take in a key and return true if the bucket has the key
func (b *bucket) BucketSearch(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.Key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//BucketDelete will take in a key and delete from the bucket
func (b *bucket) BucketDelete(k string) {
	if b.head.Key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next.Key != k {
		if previousNode.next.Key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

//Hash
func Hash(s string) int {
	sum := 0
	for _, i := range s {
		sum += int(i)
	}
	return sum % ArraySize
}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	myHashTable := Init()
	fmt.Println(myHashTable)

	fmt.Println(Hash("Rama"))

	testBucket := &bucket{}
	testBucket.BucketInsert("Randy")
	testBucket.BucketDelete("Randy")
	fmt.Println(testBucket.BucketSearch("Randy"))
	fmt.Println(testBucket.BucketSearch("barbie"))

	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY", //hashcode - 4
		"KYLE",  //hashcode - 4
		"STAN",
		"Thiyagu",
		"Kavin",
		"Kuhan",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")

	fmt.Println("Kavin", hashTable.Search("Kavin"))
	fmt.Println("Kuhan", hashTable.Search("Kuhan"))
	fmt.Println("uuuu", hashTable.Search("uuuu"))
}
