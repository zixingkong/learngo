// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// 3. After you've done, check out the solution
//    and read the comments there
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

func main() {
	// var i int
	// var i8 int8

	// CONTINUE FROM HERE....
	//    int
	var it int 
	fmt.Println(it)
	//  int8
	var i8 int8
	fmt.Println(i8)
	//    int16
	var i6 int16
	fmt.Println(i6)
	//    int32
	var i32 int32
	fmt.Println(i32)
	//    int64
	var i64 int64
	fmt.Println(i64)
	//    float32
	var f32 float32
	fmt.Println(f32)
	//    float64
	var f64 float64
	fmt.Println(f64)
	//    complex64
	var c64 complex64
	fmt.Println(c64)
	//    complex128
	var c128 complex128
	fmt.Println(c128)
	//    bool
	var bl bool
	fmt.Println(bl)
	//    string
	var s string
	fmt.Println(s)
	//    rune
	var rn rune 
	fmt.Println(rn)
	//    byte
	var bt byte 
	fmt.Println(bt)
}
