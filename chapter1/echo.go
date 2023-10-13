package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""

	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i]
	}

	fmt.Println(s)
}
