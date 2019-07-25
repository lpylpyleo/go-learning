package main

import (
  "fmt"
)

func changeArray(arr *[3]int) {
  arr[0]=999
}

func main(){
  arr1 := [3]int{1,2,3}
  changeArray(&arr1)
  fmt.Println(&arr1[0],arr1)
  arr2 := [5]int{1,2}
  slc := make([]int,3)
  slc[2]=1
  slc = append(slc,5)
  fmt.Println(len(arr2),cap(arr2))
  fmt.Println(slc,len(slc),cap(slc))
  //fmt.Println(arr1[-1:])
}