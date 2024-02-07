package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type W struct {
	b int32
	c int32
}

type student struct {
	Id   int
	Name string
}

type Person struct {
	Stu *student
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	p := &Person{}
	for i := 1; i < 11; i++ {
		go func(i int) {
			p.Stu = &student{
				Id: i,
			}
			wg.Done()
		}(i)
	}
	for j := 1; j < 11; j++ {
		go func(j int) {
			fmt.Println(p.Stu.Id)
			wg.Done()
		}(j)
	}
	wg.Wait()
	return
	// 能说说uintptr和unsafe.Pointer的区别吗？
	// unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
	// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
	// unsafe.Pointer 可以和 普通指针 进行相互转换；
	// unsafe.Pointer 可以和 uintptr 进行相互转换。

	var w *W = new(W)
	fmt.Println(w.b, w.c)
	fmt.Println(unsafe.Sizeof(w))

	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 1
	fmt.Println(*((*int)(b)))
	fmt.Println(w.b, w.c)
	fmt.Println(unsafe.Sizeof(w))
}
