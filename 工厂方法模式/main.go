package main

import factory "p5/factory"
import "fmt"

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _:=factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}