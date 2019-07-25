package main

import "fmt"

func main(){
  arr := [5]int{2,3,5,7,11}
  for index,value := range arr{
    fmt.Printf("arr[%d] = %d\n",index,value)
  }
}