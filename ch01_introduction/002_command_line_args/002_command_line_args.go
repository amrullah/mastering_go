package main

import (
	"fmt"
	"os"
)

func main() {
	var osArgs string
	separator := " "
	for i := 1; i < len(os.Args); i++ {
		if i == 1 {
			osArgs += os.Args[i]
		} else {
			osArgs += separator + os.Args[i]
		}

	}
	fmt.Println(osArgs)
}
