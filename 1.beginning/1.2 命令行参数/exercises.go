package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//exercise 1 and 3
	start := time.Now()
	sj := strings.Join(os.Args[:], "-")
	fmt.Printf("%.2fs %v\n", time.Since(start).Seconds(), sj)

	start1 := time.Now()
	var s, sep string
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = ","
		//exercise 2
		fmt.Printf("%.2fs %d %v\n", time.Since(start1).Seconds(), i, s)
	}
	fmt.Printf("%.2fs %v", time.Since(start1).Seconds(), s)
}
