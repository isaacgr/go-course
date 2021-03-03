package main

import "fmt"

type colorMap map[string]string

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := colorMap{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
	fmt.Println(colors)
}

func printMap(c colorMap) {
	for key, value := range c {
		fmt.Println("Color: ", key, " ", "Hex: ", value)
	}
	c["blue"] = "#0000ff"
}
