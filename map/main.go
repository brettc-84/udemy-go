package main

import "fmt"

func main() {
	// var colours map[string]string
	colours := make(map[string]string)
	colours = map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"black": "unknown",
	}

	colours["white"] = "#ffffff"

	// delete(colours, "red")

	// fmt.Println(colours)
	// printMap(colours)
	updateMap(colours)
	printMap(colours)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}

func updateMap(c map[string]string) {
	c["black"] = "#000000"
}
