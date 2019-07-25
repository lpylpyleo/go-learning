package main

import "fmt"

  type Person struct{
    name string
    age int
  }

func main(){
    var p = Person{"tom", 5}
    fmt.Println(p)

    var pp *Person = &p
    fmt.Println(pp)
}