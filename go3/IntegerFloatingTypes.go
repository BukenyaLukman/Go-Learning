package main

import "fmt"

func main(){

	var n complex64 = 1 + 2i
	var p complex64 = 3 + 4i
	var q complex64 = 2 + 2i


	a := 12.3
	b := 10.5

	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a/b)
	fmt.Println(a*b)

	fmt.Printf("%v, %T\n",n, n)
	fmt.Println(p + q)
}