package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to gclojure. Use ctrl-c to exit.")
	reader := NewReader(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("> ")
		obj, err := TryRead(reader)
		switch {
		case err == io.EOF:
			return
		case err != nil:
			fmt.Println("Error: ", err)
		default:
			fmt.Println(obj.ToString(false))
		}
	}
}
