package main

import "fmt"

func main() {
  var age int = 1
  var b bool = true
  var e error
  c := 36
  if(b){
    fmt.Println(b, age, e, c)
  }
  fmt.Println("----------------")
  v1, v2, v3 := true, 3.6, "45"
  fmt.Println(v1,v2,v3,&v1)
}
