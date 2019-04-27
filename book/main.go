package main

import "fmt"

// create new type book of string
type book string

// create receiver of type book
func (b book) printTitle() {
	fmt.Println(b)
}

func main() {
	var b book = "Harry Potter"
	b.printTitle()
}
