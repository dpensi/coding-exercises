package main

import (
	"fmt"

	"github.com/dpensi/coding-exercises/options"
)

func main() {
	fmt.Println(options.NewCustomer("aaa", "ciao@mail.com"))
	fmt.Println(
		options.NewCustomer("bbb", "cim",
			options.WithName("daniele", "pensiero")))

	fmt.Println(
		options.NewCustomer("ccc", "ciao@mail.com",
			options.WithName("daniele", "pensiero"),
			options.WithLoyaltyPoints(200)))

	fmt.Println(
		options.NewCustomer("ddd", "ciao@mail.com",
			options.WithName("daniele", "pensiero"),
			options.WithLoyaltyPoints(200)))
}
