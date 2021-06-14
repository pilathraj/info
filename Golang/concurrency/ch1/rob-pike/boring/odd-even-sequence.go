package main

import (
	"fmt"
	"time"
)

type Msg struct {
	n    int
	wait chan bool
}

func odd(n int) {
	if n%2 != 0 {
		fmt.Println(n)

	}

}
func even(n int) {
	if n%2 == 0 {
		fmt.Println(n)

	}
}
func main() {

	for i := 0; i < 5; i++ {
		go odd(i)
		go even(i)
	}
	time.Sleep(1 * time.Second)
}
