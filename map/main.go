package main

import "fmt"

func main() {
	// var colours map[string]string
	colours := make(map[string]string)
	colours = map[string]string{
		"red":   "#ff0000",
		"green": "##008000",
	}

	colours["white"] = "#ffffff"

	// delete(colours, "red")

	// fmt.Println(colours)
	printMap(colours)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
