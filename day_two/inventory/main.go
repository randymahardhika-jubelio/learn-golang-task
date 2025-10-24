package main

import (
	"fmt"
)

// Item models an inventory item
type Item struct {
	Name     string
	Price    float64
	Quantity int
}

// Inventory manages a collection of items
type Inventory struct {
	Items map[string]Item
}

// AddItem adds a new item or updates an existing one
func (inv *Inventory) AddItem(item Item) {
	if existing, found := inv.Items[item.Name]; found {
		existing.Price = item.Price
		existing.Quantity += item.Quantity
		inv.Items[item.Name] = existing
	} else {
		inv.Items[item.Name] = item
	}
}

// PrintInventory prints all items in the inventory
func (inv *Inventory) PrintInventory() {
	fmt.Println("Inventory Report:")
	for _, item := range inv.Items {
		fmt.Printf("Name: %s | Price: %.2f | Quantity: %d\n", item.Name, item.Price, item.Quantity)
	}
}

func main() {
	inventory := Inventory{Items: make(map[string]Item)}

	inventory.AddItem(Item{Name: "Apple", Price: 0.5, Quantity: 100})
	inventory.AddItem(Item{Name: "Banana", Price: 0.3, Quantity: 150})
	inventory.AddItem(Item{Name: "Orange", Price: 0.7, Quantity: 80})
	inventory.AddItem(Item{Name: "Apple", Price: 0.55, Quantity: 50})

	inventory.PrintInventory()
}