package main

import (
  "fmt"
  "net/http"
  "log"
  "time"
  "strconv"
)

type Student struct {
  name string
  age int
}

func (a Student) String() string {
  return a.name + ", " + strconv.Itoa(a.age)
}

var A Student

func init(){
  A = Student{name:"tom",age:20}
}

func sayHello(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Fprintf(w, time.Now().String())
  fmt.Fprintf(w, "\n")
  fmt.Fprintf(w, A.String())
}

func main(){
  http.HandleFunc("/",sayHello)
  err := http.ListenAndServe(":18888",nil)
  if err != nil {
    log.Fatal("ListenAndServe: ",err)
  }
}