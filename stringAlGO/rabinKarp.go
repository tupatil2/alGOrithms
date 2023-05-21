package main

import (
  "fmt"
)

func main() {
  s := "xoxo"
  t := "oxo"
  fmt.Println(rabinKarp(s, t))

}

// check the first occurrence of string t in s and return the index 
func rabinKarp(s string, t string) int {
  n := len(s)
  m := len(t)

  if n == 0 && m == 0 {
    return 0 
  }

  prime := (int64) (31)
  mod := (int64) (1000000007)

  tHash := (int64) (0)
  pr := (int64) (1)

  for i := 0; i < m; i++ {
    tHash = (tHash + (((int64)(t[i] - 'a')) + 1) * pr) % mod
    pr = (pr * prime) % mod
  }

  index := -1
  sHash := (int64) (0)
  pr = (int64)(1)

  for i := n-1; i >= 0; i-- {
    sHash = ((sHash * prime) % mod + ((int64)(s[i]-'a')) + 1) % mod

    if i + m >= n {
      pr = (pr * prime) % mod
    } else {
      sHash = (sHash - ((((int64)(s[i+m]-'a'))+1) * pr) % mod + mod) % mod 
    }

    if sHash == tHash {
      index = i
    }
  }

  return index
}
