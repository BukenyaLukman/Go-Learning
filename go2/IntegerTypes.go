package main

import "fmt"

func main(){
	a := 10
	b := 3

	fmt.Printf("%v\n", a+b)
	fmt.Printf("%v\n", a-b)
	fmt.Printf("%v\n", a*b)
	fmt.Printf("%v\n", a/b)
	fmt.Printf("%v\n", a%b)
	fmt.Printf("%v\n", a^b)
	fmt.Printf("%v\n", a&^b)
	fmt.Printf("%v\n", a|b)
}