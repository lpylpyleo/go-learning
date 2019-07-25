package main

import "fmt"

func main(){
  var a = [...]int {1,2,3}

  for i:=0;i<len(a);i++ {
    fmt.Printf("a[%d] = %d\n",i,a[i])
  }
}