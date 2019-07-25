package main

import "fmt"

const (
  a = iota
  b
  c
)

const (
  d=iota
  e
  f='a'
  g
)

func main(){
  fmt.Println(a,b,c,d,e,f,g)
}