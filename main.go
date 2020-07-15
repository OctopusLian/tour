package main

import (
	"log"

	"github.com/programming-book-practice/tour/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
