package main

import (
	
	"time"
	"fmt"
)

func tick(d time.Duration) <-chan int{
  c := make(chan int)
    go func(){
      time.Sleep(d)
      c <- 1
    }()
   return c
}
func main() {
   for i :=0; i<24; i++{
   ch := tick(100 * time.Millisecond)
   fmt.Println(i, ": ", <-ch)
   }	
}
