package main

import "fmt"

func main() {
	george := &Grandpa{
		name: "George Joestar",
	}

	jonathan := &Father{
		name:     "Jonathan Joestar",
		ancestor: george,
	}

	joseph := &Son{
		name:     "Joseph Joestar",
		ancestor: jonathan,
	}

	fmt.Print(joseph.getAncestry())
}
