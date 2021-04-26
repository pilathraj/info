package main

import "fmt"

func main(){
  l, err := net.Listen("tcp", ":5000")
  if err != nil{
    panic(err)
  }
}
