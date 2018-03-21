package main

import "fmt"

func main()  {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var b1, c1 float64 = 1.1, 2.1
	fmt.Println(b1, c1)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println("e", e)

	//The := syntax is shorthand for declaring and initializing a variable
	f := "short"
	fmt.Println(f)

	const n = 500000000
	fmt.Println(n)

	//n = n*5
	//cannot assign to n

	const n2 = 3e20 / n
	fmt.Println(n2)
	fmt.Println(int64(n2))

}