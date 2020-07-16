package main

import "fmt"

func main(){

	s := "Hello World"
	s1 := " Good morning"
	s3 := s+s1
	fmt.Printf("%v\n",s3)

	p := "Hello world"
	b := []byte(p)

	fmt.Printf("%v, %T\n",b, b)

	r := 'm'
	fmt.Printf("%v, %T",r,r)

	// Converting Strings into stream of bytes
	// Slices of bytes


}
