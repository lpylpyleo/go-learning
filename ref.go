package main

import "fmt"

func main(){
  a, b, c :=1, 2, 3
  var r1 = &a
  var r2 = r1
  fmt.Println(r1, r2)
  r1 = &b
  fmt.Println(r1, r2, c)
}