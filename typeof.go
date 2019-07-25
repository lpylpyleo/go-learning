package main

import (
   "reflect"
   "fmt"
)


func main(){
  var a []int
  var b [5]int
  fmt.Println(reflect.TypeOf(a).String())
  fmt.Println(reflect.TypeOf(b).String())
}