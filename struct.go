package main

type Player struct {
	Inventory       []string
	CurrentLocation *Location
}

type Location struct {
	Name        string
	DescMove    string
	DescLook    DescLook
	Items       []string
	Connections map[string]*Location
}

type DescLook struct {
	WithItems string
	NoItems   string
}
