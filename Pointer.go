package main

import "fmt"

func main(){
  var p, q *int
  a := 1
  p = &a
  fmt.Println(p, q)
}