package main

import (
	"fmt"
)

func boring(s string) <- chan string {
 ch := make(chan string)
 go func(){
   for i:=0;;i++{
    ch <- fmt.Sprintf("%s %d", s, i)
   }
 }()
 return ch

} 

func main() {
	fmt.Println("Hello, playground")
	c := boring("boring") // boring function return the channel.
	fmt.Println(c, len(c))
	for i:=0; i<5; i++{
	  fmt.Println(<-c)
	}
	fmt.Println("Program exits")
	
}
