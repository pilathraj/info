package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func(){
	  ch<-42
	  ch<-22
	}()
	fmt.Println(<-ch) //42
	fmt.Println(<-ch) // 22
	//fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock!
	fmt.Println("Main END")
	
}
