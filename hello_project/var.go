package main

import ("fmt"
	"unicode/utf8")

var a = 3
func main() {
	//var b int

//	c := 3
	//var s1 string="Hello World"
	var s2 string=`안녕하세요
	Hell World`
	fmt.Println(s2)
	s2 = "한글"
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2))

	var s1 = "Hello"
	s2 = "Hello"
	var s3 = "World"

	fmt.Println(s1==s2)
	fmt.Println(s1+s3)
	
	//	var V1 bool = true
	//var V2 bool = false

	//const a = 10
	
	const (
		a = 0
		b = 1
		c = 2
		d = 3
		f = 4
	)

	/*
	const (
		a=itoa
		b
		c
		d
		e
	)
	*/

}

