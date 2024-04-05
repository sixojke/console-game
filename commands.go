package main

import (
	"fmt"
	"strings"
)

func (p *Player) Look() {
	var look string
	if len(p.CurrentLocation.Items) == 0 {
		look = p.CurrentLocation.DescLook.NoItems + "можно пройти - " + strings.Join(moveLocations(p.CurrentLocation.Connections), ", ") + "."
	} else {
		look = p.CurrentLocation.DescLook.WithItems + strings.Join(p.CurrentLocation.Items, ", ") + ". " + "можно пройти - " +
			strings.Join(moveLocations(p.CurrentLocation.Connections), ", ") + "."
	}
	fmt.Println(look)
}

func (p *Player) TakeItem(itemName string) {
	for i, item := range p.CurrentLocation.Items {
		if item == itemName {
			p.Inventory = append(p.Inventory, itemName)
			p.CurrentLocation.Items = append(p.CurrentLocation.Items[:i], p.CurrentLocation.Items[i+1:]...)
			fmt.Printf("Предмет добавлен в инвентарь: %s\n", itemName)
			return
		}
	}
	fmt.Printf("Предмет %s не найден в локации.\n", itemName)
}

func (p *Player) MoveTo(locationName string) {
	nextLocation, ok := p.CurrentLocation.Connections[locationName]
	if ok {
		p.CurrentLocation = nextLocation
		fmt.Println(p.CurrentLocation.DescMove + "можно пройти - " + strings.Join(moveLocations(p.CurrentLocation.Connections), ", ") + ".")
	} else {
		fmt.Println("Невозможно перейти в указанную локацию.")
	}
}

func moveLocations(connections map[string]*Location) []string {
	var moveLocations []string
	for location := range connections {
		moveLocations = append(moveLocations, location)
	}
	return moveLocations
}

func (p *Player) ListCommands() {
	fmt.Println("Доступные команды:")
	fmt.Println("- look: осмотреться вокруг")
	fmt.Println("- go <название локации>: переместиться в другую локацию")
	fmt.Println("- take <предмет>: взять предмет из локации")
	fmt.Println("- inventory: просмотреть инвентарь")
	fmt.Println("- help: показать список доступных команд")
}
