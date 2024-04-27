package main

import "fmt"

var testing string = "Go is weird"

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hello World"
	fmt.Println(y)

	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)

	var a string
	a = "first"
	fmt.Println(a)
	a = a + "second"
	fmt.Println(a)

	var t string = "hi"
	var n string = "hi"
	fmt.Println(t == n)

	c := "hello hi"

	fmt.Println(c)
	var i = "Hi"
	fmt.Println(i)

	d := 5
	fmt.Println(d)

	name := "Max"
	fmt.Println("My dog's name is", name)

	fmt.Println(testing)
}

func f() {
	fmt.Println(testing)
}
