package main

import (
  "fmt"
)

func main() {
  stack := Stack{}
  stack.Push(1)
  stack.Push(2)
  stack.Push(3)
  stack.Push(4)
  stack.Push(5)
  stack.Push(6)
  stack.Push(7)

  fmt.Println("size: ", stack.Size())
  fmt.Println("peek: ", stack.Peek())

  element := stack.Pop()
  fmt.Println("popped element: ", element)

  fmt.Println("size: ", stack.Size())

  for stack.Size() != 0 {
    fmt.Println("element is ", stack.Pop())
  }
} 

type Stack struct {
  elements []interface{}
}

func (s *Stack) Push(item interface{}) {
  s.elements = append(s.elements, item)
}

func (s *Stack) Pop() interface{} {
  len := len(s.elements)
  if len == 0 {
    return nil
  }
  item := s.elements[len-1]
  s.elements = s.elements[0:len-1]
  return item 
}

func (s *Stack) Peek() interface{} {
  len := len(s.elements)
  if len == 0 {
    return nil
  }
  return s.elements[len-1] 
}

func (s *Stack) Size() int {
  return len(s.elements)
}
