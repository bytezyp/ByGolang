package main

func main12() {
	unbuffered := make(chan int)
	//a := <- unbuffered // 阻塞
	unbuffered <- 1 // 阻塞
	//fmt.Println(a)
}
