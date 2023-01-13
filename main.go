package main

import (
	interfaces "abstract_factory/interfaces"
	factory "abstract_factory/factory"
	"fmt"
)

func main() {
	adidasFactory, _ := factory.GetSportsFactory("adidas")
	nikeFactory, _ := factory.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()
	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s interfaces.IShoe) {
	fmt.Printf("Shoe logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Shoe size: %d", s.GetSize())
	fmt.Println()
}

func printShortDetails(s interfaces.IShort) {
	fmt.Printf("Short logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Short size: %d", s.GetSize())
	fmt.Println()
}
