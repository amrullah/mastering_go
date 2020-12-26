package main

import (
	"fmt"
	"os"
)

func main() {
	osArgs, separator := "", ""

	for _, arg := range os.Args[1:] {
		osArgs += separator + arg
		separator = " "
	}

	//osArgs := strings.Join(os.Args[1:], " ") // alternative to everything above

	fmt.Println(osArgs)
}
