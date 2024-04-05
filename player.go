package main

import "fmt"

func CreatePlayer(startLocation string) (*Player, error) {
	for _, location := range Locations {
		if location.Name == startLocation {
			return &Player{
				Inventory:       []string{},
				CurrentLocation: location,
			}, nil
		}
	}

	return nil, fmt.Errorf("start location not found")
}
