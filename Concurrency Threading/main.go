package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU Cores:", runtime.NumCPU())
}

//Main is also Go routine 
// No Goroutine waits to other  Go routine to complete Execution 
//if main Exists, the entire Application exits 
