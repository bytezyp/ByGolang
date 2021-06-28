package slicestacks

import (
	"errors"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

func New(values ...interface{}) *Stack {
	stack := Stack{make([]interface{}, 0, len(values))}
	//fmt.Println(values)
	stack.Push(values...)
	return &stack
}

// 定义栈的结构
type Stack struct {
	array []interface{}
}

// Push 添加元素到栈中
func (s *Stack) Push(values ...interface{}) {
	s.array = append(s.array, values...)
}

//Size 判断栈的大小
func (s *Stack) Size() int {
	return len(s.array)
}

// IsEmpty 判断是否为空
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Clear 重置栈
func (s *Stack) Clear() {
	s.array = nil
}
