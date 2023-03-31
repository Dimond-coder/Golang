package main

import (
	"fmt"
)

// // typed constant
// const a int = 10
// // untyped constant
// const PI float32 = 3.14

const (
	a         = 1
	b string  = "10"
	c float32 = 12.299
	d string  = "passion"
)

func main() {
	//var arr1 = [...]string{"hello", "hi", "anhmo", "daunhieuthem", "bietdau"}
	var arr2 = [6]string{"heloo", "hi", "daunhieuthem", "bietdau"}
	var arr3 = [...]string{"heloo", "hi", "daunhieuthem", "bietdau"}
	var arr = []int{1: 10, 3: 60}

	//fmt.Print("\n", arr1[], arr1[], arr1[])
	fmt.Println()
	fmt.Println(arr2)
	fmt.Println()
	fmt.Println(arr3)
	fmt.Println(arr)

	// fmt.Print(a, PI)
	// fmt.Println()
	// fmt.Print(a, PI)
	// fmt.Println()
	// fmt.Printf("%v %T", a, PI)
	// //Declared inside
	// const b = 10
	// fmt.Println("\n", b)
	//fmt.Printf("%+d\n %s\n %#v\n %q", a, b, c, d)
	//fmt.Println(a, b, c, d)

	//fmt.Println()
	//fib()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	// fmt.Println("Hello world!")

	// // tao bien va gan gia tri cho bien
	// var x = 9
	// var y = 10
	// var sum = x + y
	// fmt.Print(sum)

	//Tao gia tri cho bien cach nhanh hon
	// x := 10
	// y := 9
	// sum := x + y
	// fmt.Println(sum)

	//su dụng if-else if -else
	// var a = 9
	// if a > 20 {
	// 	fmt.Print("a more than 20!")
	// } else if a == 10 {
	// 	fmt.Print("a equal to 10")
	// } else {
	// 	fmt.Print("e smaller than 20 and a not equal to 10!")
	// }

	//uses Array basic
	// var a [5]int
	// fmt.Print(a)

	// //uses Array shorted
	// a := []int{5, 4, 3, 4, 5}
	// b := []int{10, 20}
	// a = append(a, b...)
	// fmt.Println(a)

}
