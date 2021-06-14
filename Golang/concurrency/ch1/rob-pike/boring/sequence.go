package main

import (
	"fmt"
	//"time"
	//"math/rand"
)

type Msg struct {
	Str  string
	Wait chan bool
}

func boring(msg string) <-chan Msg {

	ch := make(chan Msg)
	waitFor := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			ch <- Msg{fmt.Sprintf("%s %d", msg, i), waitFor}
			//time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitFor
		}
	}()
	return ch
}

func fanin(input1, input2 <-chan Msg) <-chan Msg {
	c := make(chan Msg)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanin(boring("pilathraj"), boring("ragu"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.Str)
		msg2 := <-c
		fmt.Println(msg2.Str)
		msg1.Wait <- true
		msg2.Wait <- true
	}
}
