package main

import (
	"fmt"
	"sync/atomic"
)

func demo() int {
	return 1111

	defer func() {
		fmt.Println(3333)
	}()
	return 2222
}

func main() {
	fmt.Println(demo())
	//a := false && false || true || true
	//fmt.Println(a)
	//return
	//array := [][]int{{1, 2}, {3}, {4, 5}}
	//result := sortAll(array)
	//fmt.Println(result)
	var a atomic.Value
	a.CompareAndSwap()
}

func sortAll(array [][]int) [][]int {
	result := make([][]int, 0, len(array))
	_sortAll(array, 0, nil, &result)
	return result
}

func _sortAll(array [][]int, index int, list []int, result *[][]int) {
	fmt.Println(index, array, list)
	if index == len(array) {
		newList := make([]int, len(list))
		copy(newList, list)
		*result = append(*result, newList)
		return
	}
	curList := array[index]
	for i := 0; i < len(curList); i++ {
		list = append(list, curList[i])
		_sortAll(array, index+1, list, result)
		//list = list[:len(list)-1]
	}
}
