package main

import (
  "fmt"
)

func main() {

  arr := []int{9,8,7,6,5,4,3,2,1}

  ans := mergeSort(arr, 0, len(arr) - 1)
  fmt.Println(ans)
}

func mergeSort(arr []int, start int, end int) []int {
  if(start > end) {
    return []int{}
  }
  if(start == end) {
    return []int{arr[start]} 
  }

  mid := start + (end - start) / 2

  left := mergeSort(arr, start, mid)
  right := mergeSort(arr, mid + 1 , end)
  
  return merge(left, right)
}

func merge(left []int, right []int) []int {
  n := len(left)
  m := len(right)

  var merge []int
  leftIndex := 0
  rightIndex := 0

  for leftIndex < n || rightIndex < m {
    if leftIndex == n {
      merge = append(merge, right[rightIndex])
      rightIndex++
    } else if rightIndex == m {
      merge = append(merge, left[leftIndex])
      leftIndex++
    } else if left[leftIndex] <= right[rightIndex]{
      merge = append(merge, left[leftIndex])
      leftIndex++
    } else {
      merge = append(merge, right[rightIndex])
      rightIndex++
    }
  }

  return merge
}
