package main

import(
  "fmt"
)

func main() {
  N := 100
  prime := sieve(N)
  for i := 0; i <= N; i++ {
    if prime[i] {
      fmt.Printf("%v is prime \n", i)
    }
  }
}

func sieve(n int) []bool {
  prime := make([]bool, n+1)
  for i := 2; i <= n; i++ {
    prime[i] = true
  }

  for i := 2; i <= n; i++ {
    if prime[i] {
      for j := 2 * i; j <= n; j += i {
        prime[j] = false
      }
    }
  }

  return prime
}
