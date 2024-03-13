package mypackage

import (
	"fmt"
	"runtime"
)


func Version() string {
	fmt.Println("Hello mypackage")
    return runtime.Version()
}