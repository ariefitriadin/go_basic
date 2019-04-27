package main

import "fmt"

type laptopSize float64

func main() {
	// initialize variable size with type of laptopSize
	var size laptopSize = 3.5
	fmt.Println(size.getSizeOfLaptop())
}

func (ls laptopSize) getSizeOfLaptop() laptopSize {
	return ls
}
