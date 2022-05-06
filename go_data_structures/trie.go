package main

import "fmt"

type TrieNode struct {
	children map[string]TrieNode
}

func initializeTrieNode() TrieNode {
	newTrieNode := TrieNode{}
	newTrieNode.children = map[string]TrieNode{}
	//newTrieNode.children["t"] = initializeTrieNode()
	return newTrieNode
}

type Trie struct {
	root TrieNode
}

func initializeTrie() Trie {
	newTrie := Trie{}
	newTrie.root = initializeTrieNode()
	//newTrie.root.children["t"] = initializeTrieNode()
	return newTrie
}

// func (t *Trie) search() interface{} {

// }

func (t *Trie) insert(word string) {
	currentNode := t.root

	for _, elem := range word {
		_, ok := currentNode.children[string(elem)]
		if ok {
			currentNode = currentNode.children[string(elem)]
		} else {
			newNode := initializeTrieNode()
			currentNode.children[string(elem)] = newNode
			currentNode = newNode
		}
	}

	currentNode.children["*"] = initializeTrieNode()
}

func (t *Trie) search(word string) bool {
	currentNode := t.root

	for _, elem := range word {
		_, ok := currentNode.children[string(elem)]
		if ok {
			currentNode = currentNode.children[string(elem)]
		} else {
			return false
		}
	}

	return true
}

func (t *Trie) collectAllWords(currentNode TrieNode, word string, words *[]string) []string {

	for elem, child := range currentNode.children {
		fmt.Println(elem)
		if elem == "*" {
			*words = append(*words, word)
			fmt.Println(words)
		} else {
			updatedWord := word + string(elem)
			t.collectAllWords(child, updatedWord, words)
		}
	}

	return *words
}

func main() {
	newTrie := initializeTrie()
	fmt.Println(newTrie.root.children)
	newTrie.insert("test")
	fmt.Println(newTrie)
	newTrie.insert("tesp")
	newTrie.insert("pelt")
	fmt.Println(newTrie)
	fmt.Println(newTrie.search("test"))
	words := []string{}
	fmt.Println(newTrie.collectAllWords(newTrie.root, "", &words))
}
