package main

import (
	"fmt"
)

func main() {
	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// 声明两个含有byte的slice
	var a, b []byte

	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]

	// b是数组ar的另一个slice
	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]

  c := a[:]

  a[0] = 0

  fmt.Println(ar)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
