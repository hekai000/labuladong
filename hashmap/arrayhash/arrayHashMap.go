package arrayhash

import (
	"fmt"
	"math/rand"
)

type Node struct {
	key int
	val int
}

type MyArrayHashMap struct {
	arr []Node
	m   map[int]int
}

func NewMyArrayHashMap() *MyArrayHashMap {
	return &MyArrayHashMap{
		arr: make([]Node, 0),
		m:   make(map[int]int),
	}
}

func (this *MyArrayHashMap) Get(key int) int {
	if _, ok := this.m[key]; ok {
		return this.arr[this.m[key]].val
	}
	return -1
}

func (this *MyArrayHashMap) Put(key int, val int) {
	if this.containsKey(key) {
		this.arr[this.m[key]].val = val
		return
	}
	this.arr = append(this.arr, Node{key, val})
	this.m[key] = len(this.arr) - 1
}

func (this *MyArrayHashMap) containsKey(key int) bool {
	_, ok := this.m[key]
	return ok
}

func (this *MyArrayHashMap) size() int {
	return len(this.m)
}

func (this *MyArrayHashMap) randomKey() int {
	n := len(this.arr)
	randomIndex := rand.Intn(n)
	return this.arr[randomIndex].key
}

func (this *MyArrayHashMap) Remove(key int) {
	if _, ok := this.m[key]; !ok {
		return
	}
	index := this.m[key]
	node := this.arr[index]

	e := this.arr[len(this.arr)-1]
	this.arr[index] = e
	this.arr[len(this.arr)-1] = node
	this.m[e.key] = index
	this.arr = this.arr[:len(this.arr)-1]
	delete(this.m, node.key)
}

func main() {
	arrMap := NewMyArrayHashMap()
	arrMap.Put(1, 1)
	arrMap.Put(2, 2)
	arrMap.Put(3, 3)
	arrMap.Put(4, 4)
	arrMap.Put(5, 5)

	fmt.Println(arrMap.Get(1)) // 1
	fmt.Println(arrMap.randomKey())

	arrMap.Remove(4)
	fmt.Println(arrMap.randomKey())
	fmt.Println(arrMap.randomKey())
}
