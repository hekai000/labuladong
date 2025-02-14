package main

import "fmt"

type Node struct {
	key  string
	val  int
	prev *Node
	next *Node
}
type MyLinkedHashMap struct {
	head *Node
	tail *Node
	m    map[string]*Node
}

func Constructor() MyLinkedHashMap {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return MyLinkedHashMap{
		head: head,
		tail: tail,
		m:    make(map[string]*Node),
	}
}

func (this *MyLinkedHashMap) Get(key string) int {
	if node, ok := this.m[key]; ok {
		return node.val
	}
	return -1
}

func (this *MyLinkedHashMap) Put(key string, val int) {
	if _, ok := this.m[key]; !ok {
		node := &Node{
			key: key,
			val: val,
		}
		this.addLastNode(node)
		this.m[key] = node
		return
	}
	this.m[key].val = val

}

func (this *MyLinkedHashMap) addLastNode(node *Node) {
	temp := this.tail.prev
	node.next = this.tail
	node.prev = temp

	temp.next = node
	this.tail.prev = node

}

func (this *MyLinkedHashMap) Remove(key string) {
	if _, ok := this.m[key]; !ok {
		return
	}
	node := this.m[key]
	delete(this.m, key)
	this.removeNode(node)
}

func (this *MyLinkedHashMap) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	node.prev = nil
	node.next = nil
}

func (this *MyLinkedHashMap) ContainsKey(key string) bool {
	_, ok := this.m[key]
	return ok
}

func (this *MyLinkedHashMap) Keys() []string {
	keyList := make([]string, 0)
	for p := this.head.next; p != this.tail; p = p.next {
		keyList = append(keyList, p.key)
	}
	return keyList
}

func main() {
	myMap := Constructor()
	myMap.Put("a", 1)
	myMap.Put("b", 2)
	myMap.Put("c", 3)
	myMap.Put("d", 4)
	myMap.Put("e", 5)

	// output: a b c d e
	fmt.Println(myMap.Keys())

	myMap.Remove("c")

	// output: a b d e
	fmt.Println(myMap.Keys())
}
