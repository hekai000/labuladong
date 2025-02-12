package main

import "fmt"

func main() {
	//定义静态数组
	var arr [10]int
	arr[0] = 1
	arr[2] = 2

	a := arr[0]
	fmt.Println(a)
}

func arrayAppendBack() {
	var arr [10]int
	for i := 0; i < 4; i++ {
		arr[i] = i
	}
	arr[4] = 4
	arr[5] = 5
}

func arrayAppendMiddle() {
	var arr [10]int
	for i := 0; i < 4; i++ {
		arr[i] = i
	}
	for i := 4; i > 2; i-- {
		arr[i] = arr[i-1]
	}
	arr[2] = 666
	fmt.Println(arr)
}

func extendArray() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	newArr := make([]int, 20)
	for i := 0; i < 10; i++ {
		newArr[i] = arr[i]
	}
	newArr[10] = 10
}

func deleteBack() {
	var arr [10]int
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	arr[4] = -1
}

func deleteMiddle() {
	var arr [10]int
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	for i := 1; i < 5; i++ {
		arr[i] = arr[i+1]
	}
	arr[4] = -1
}

func dynamicArray() {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	//在索引为2的位置插入666
	arr = append(arr[:2], append([]int{666}, arr[2:]...)...)

	//在头部插入元素
	arr = append([]int{-1}, arr...)

	//删除末尾元素
	arr = arr[:len(arr)-1]

	//删除索引为2的元素
	arr = append(arr[:2], arr[3:]...)

	a := arr[0]
	arr[0] = 100
	fmt.Println(a)

	index := -1
	for i, v := range arr {
		if v == 666 {
			index = i
			break
		}
	}
	fmt.Println(index)
}
