package main

import (
	"fmt"
	"time"
)

func producer(in chan int,  d time.Duration){
  for {
    in <-1
    time.Sleep(d)
  }
}

func outer(out chan int){
  for {
   fmt.Println(<-out)
  }
}

func main() {
	in, out := make(chan int), make(chan int)
	go producer(in, 100 * time.Millisecond)
	go producer(in, 100 * time.Millisecond)
	go outer(out)
	for { out<- <- in}
	
	
}
