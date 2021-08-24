package main

import "fmt"

//const Albhabets = 26

//Node represents each node in a trie
type Node struct {
	children [26]*Node
	isEnd    bool
}

//Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

//InitTrie will create a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert will take a word and add it in the Trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//search will take in a word and return true if that word is insert in the Trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()
	myTrie.Insert("orc")
	fmt.Println(myTrie.Search("orc"))

	toAdd := []string{
		"orago",
		"apple",
		"erica",
		"oregano",
		"ball",
		"orea",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.Search("ball"))
	fmt.Println(myTrie.Search("uuuu"))
}
