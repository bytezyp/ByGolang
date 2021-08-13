package main

import "fmt"

type Stu struct {
	Ab string
	Bb int
}

func New(str string) *Stu {
	Ss := &Stu{
		Ab: str,
	}
	return Ss
}
func main()  {
	Ss := New("1223")
	fmt.Printf("%p \n", Ss)
	// 验证*指针复制的方式
	*Ss = Stu{
		Ab: "344",
		Bb: 22,
	}
	fmt.Printf("%p", Ss)
	fmt.Println(Ss)
}
