package main

import "fmt"

func dummy() (n int, err string) {
   return 1, "Error message"
}

func main() {
	fmt.Println("Hello, This is GO here!")
        n, err := dummy()
        fmt.Printf("n: %v, err: %v", n, err)
}
