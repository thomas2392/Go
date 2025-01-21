package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(colors)

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is", hex)
	}
}
