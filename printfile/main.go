package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1] //go run main.go myfile.txt
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
