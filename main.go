package main

import (
	"fmt"
	"github.com/dpensi/coding-exercises/graph"
)

func main() {
	fmt.Println("hello world!")
	start := graph.GetGraph1()
	result := graph.Times(start)
	for _, t := range result.Values() {
		fmt.Printf("%v %v %v\n", t.Value, t.Start(), t.Finish())
	}
}
