package main

import "fmt"

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return ch
}

func main() {
	c1 := boring("pilathraj")
	c2 := boring("ragu")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
	fmt.Println("Both boring functions exist")
}
