package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	var userNumber int
	var v [30]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		v[i] = rand.Intn(100)
	  }
	  count := 0
	  for _, n := range v{
		  if n == userNumber{
			  count ++
		  }
	  }
	  fmt.Println("Vitya, print number from 0 to 100")
	  fmt.Scan(&userNumber)
	  fmt.Println(v)
	  fmt.Printf("Frequency of %d in given array is %d.\n", userNumber, count)
}

