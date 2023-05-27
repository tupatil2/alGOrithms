package main

import (
  "fmt"
)

func main() {
  queue := Queue{}
  queue.Enqueue(1)
  queue.Enqueue(2)
  queue.Enqueue(3)
  queue.Enqueue(4)
  queue.Enqueue(5)
  queue.Enqueue(6)
  queue.Enqueue(7)

  fmt.Println("size: ", queue.Size())
  fmt.Println("peek: ", queue.Peek())
  
  firstElement := queue.Dequeue()
  fmt.Println("firstElement is ", firstElement)

  fmt.Println("size: ", queue.Size())

  for queue.Size() != 0 {
    fmt.Println("element is ", queue.Dequeue())
  }

}

type Queue struct {
  elements []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
  q.elements = append(q.elements, item)
}

func (q * Queue) Dequeue() interface{} {
  if len(q.elements) == 0 {
    return nil
  }
  item := q.elements[0]
  q.elements = q.elements[1:]
  return item
}

func (q * Queue) Peek() interface{} {
  if len(q.elements) == 0 {
    return nil
  }
  return q.elements[0]
}



func (q *Queue) Size() int {
  return len(q.elements)
}
