package main

import "fmt"

type mm struct {
	isthere bool
}

type mock struct {
	wow mm
}

func main() {
	ptest := mock{
		wow: mm{
			isthere: true,
		},
	}

	stest := mock{}

	if !stest.wow.isthere {
		fmt.Printf("Is false")
	}

	fmt.Println("This exists?", ptest.wow, "\n This does not exists", stest.wow)
}
