package main

import (
  "fmt"
  "net/http"
  "log"
)

func sayHello(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Fprintf(w, "Hello")
}

func main(){
  http.HandleFunc("/",sayHello)
  err := http.ListenAndServe(":18888",nil)
  if err != nil {
    log.Fatal("ListenAndServe: ",err)
  }
}