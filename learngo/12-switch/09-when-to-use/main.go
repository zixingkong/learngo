// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Gimme a month name.")
	// 	return
	// }

	// m := os.Args[1]

	// if m == "Dec" || m == "Jan" || m == "Feb" {
	// 	fmt.Println("Winter")
	// } else if m == "Mar" || m == "Apr" || m == "May" {
	// 	fmt.Println("Spring")
	// } else if m == "Jun" || m == "Jul" || m == "Aug" {
	// 	fmt.Println("Summer")
	// } else if m == "Sep" || m == "Oct" || m == "Nov" {
	// 	fmt.Println("Fall")
	// } else {
	// 	fmt.Println("ERROR: not a month.")
	// }

	// switch m := os.Args[1]; m {
	// case "Dec", "Jan", "Feb":
	// 	fmt.Println("Winter")
	// case "Mar", "Apr", "May":
	// 	fmt.Println("Spring")
	// case "Jun", "Jul", "Aug":
	// 	fmt.Println("Summer")
	// case "Sep", "Oct", "Nov":
	// 	fmt.Println("Fall")
	// default:
	// 	fmt.Printf("%q is not a month.\n", m)
	// }
	richter := 2.5

	switch r := richter; {
	case r < 2:
		fallthrough
	case r >= 2 && r < 3:
		fallthrough
	case r >= 5 && r < 6:
		fmt.Println("Not important")
	case r >= 6 && r < 7:
		fallthrough
	case r >= 7:
		fmt.Println("Destructive")
	}
}
