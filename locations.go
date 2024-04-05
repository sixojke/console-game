package main

var Locations = InitLocations()

func InitLocations() []*Location {
	kitchen := kitchen()
	corridor := corridor()
	bedroom := bedroom()
	street := street()

	initLocationsConnections(kitchen, corridor, bedroom, street)

	var locations []*Location
	locations = append(locations, kitchen, corridor, bedroom, street)

	return locations
}

func initLocationsConnections(kitchen, corridor, bedroom, street *Location) {
	// Локации связанные с кухней
	kitchen.Connections[corridor.Name] = corridor

	// Локации связанные с коридором
	corridor.Connections[kitchen.Name] = kitchen
	corridor.Connections[bedroom.Name] = bedroom
	corridor.Connections[street.Name] = street

	// Локации связанные с комнатой
	bedroom.Connections[corridor.Name] = corridor

	// Локации связанные с улицей
	street.Connections[corridor.Name] = corridor
}

func kitchen() *Location {
	return &Location{
		Name:     "кухня",
		DescMove: "Кухня, ничего интересного. ",
		DescLook: DescLook{
			WithItems: "На столе: ",
			NoItems:   "Ты находишься на кухне, надо собрать рюкзак и идти в универ. ",
		},
		Items:       []string{},
		Connections: make(map[string]*Location),
	}
}

func corridor() *Location {
	return &Location{
		Name:     "коридор",
		DescMove: "Ничего интересного. ",
		DescLook: DescLook{
			WithItems: "",
			NoItems:   "",
		},
		Items:       []string{},
		Connections: make(map[string]*Location),
	}
}

func bedroom() *Location {
	return &Location{
		Name:     "комната",
		DescMove: "Ты в своей комнате. ",
		DescLook: DescLook{
			WithItems: "На столе: ",
			NoItems:   "Пустая комната. ",
		},
		Items:       []string{"ключи", "конспекты", "рюкзак"},
		Connections: make(map[string]*Location),
	}
}

func street() *Location {
	return &Location{
		Name:     "улица",
		DescMove: "На улице весна. ",
		DescLook: DescLook{
			WithItems: "",
			NoItems:   "",
		},
		Items:       []string{},
		Connections: make(map[string]*Location),
	}
}
