package main

import "fmt"
import flyweight "p11/flyweight"

func main() {
	dressFactoryInstance := flyweight.GetDressFactorySingleInstance()
	dressFactoryInstance.GetDressByType(flyweight.TerroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.TerroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.TerroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.TerroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.TerroristDressType)

	dressFactoryInstance.GetDressByType(flyweight.CounterTerrroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.CounterTerrroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.CounterTerrroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.CounterTerrroristDressType)
	dressFactoryInstance.GetDressByType(flyweight.CounterTerrroristDressType)

    for dressType, dress := range dressFactoryInstance.DressMap {
        fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
    }
}