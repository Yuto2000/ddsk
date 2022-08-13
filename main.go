package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
  var counter int
  for {
    if ddsk() {
      counter ++
    } else {
      counter = 0
    }
    if counter > 2 {
      fmt.Print("ラブ注入!")
      break
    }
  }
}

func ddsk() bool {
  ddskStr := []string{"ドド", "スコ"}
  success := []int{0, 1, 1, 1}
  rand.Seed(time.Now().UnixNano())
  var count int
  for i:= 0; i < len(success); i++ {
    rand := rand.Intn(2)
    fmt.Print(ddskStr[rand])
    if (rand == success[i]) {
      count ++
    }
  }
  return count > 3
}
