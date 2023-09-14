package main

type Marker struct {
	Name        string
	Longitude   string
	Latitude    string
	Description string
	Type        string
	Tags        []Tags
}

type Tags struct {
	Name string
}
