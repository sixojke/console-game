package main

import (
	"fmt"
	"log"
)

func main() {
	player, err := CreatePlayer("кухня")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Добро пожаловать в игру!")
	player.ListCommands()

	for {
		fmt.Print("Введите команду: ")
		var command, target string
		fmt.Scanln(&command, &target)

		switch command {
		case "look":
			player.Look()
		case "go":
			player.MoveTo(target)
		case "take":
			player.TakeItem(target)
		case "inventory":
			fmt.Println("Инвентарь игрока:", player.Inventory)
		case "help":
			player.ListCommands()
		default:
			fmt.Println("Неверная команда. Введите 'help' для списка команд.")
		}
	}
}
