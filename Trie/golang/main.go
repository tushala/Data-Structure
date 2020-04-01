package main

import "fmt"

const MAXCAP = 26

type Trie struct {
	next   map[rune]*Trie
	isWord bool
}

func New() Trie {
	root := new(Trie)
	root.next = make(map[rune]*Trie, MAXCAP)
	root.isWord = false
	return *root
}

func (self *Trie) Size() int {
	return len(self.next)
}
func (self *Trie) IsEmpty() bool {
	return self.Size() == 0
}
func (self *Trie) Clear() {
	self.next = make(map[rune]*Trie, MAXCAP)
}

func (self *Trie) Add(word string) {
	for _, v := range word {
		if _, ok := self.next[v]; !ok {
			t := new(Trie)
			t.next = make(map[rune]*Trie, MAXCAP)
			t.isWord = false
			self.next[v] = t
		}
		self = self.next[v]
	}
	self.isWord = true
}

func (self *Trie) Search(word string) bool{
	for _, v := range word{
		_, ok := self.next[v]
		if !ok{
			return false
		}else{
			self = self.next[v]
		}
	}
	return self.isWord
}

func (self *Trie) StartsWith(s rune) bool {
	_, ok := self.next[s]
	return ok
}

func main() {
	trie := New()
	trie.Add("土沙拉")
	trie.Add("土土沙拉")
	fmt.Println(trie.Search("土沙拉"))
	fmt.Println(trie.Search("水沙拉"))
	fmt.Println(trie.StartsWith('水'))
	fmt.Println(trie.StartsWith('土'))
}