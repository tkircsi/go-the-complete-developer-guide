package main

import "fmt"

func main() {

	// Declaration
	// var colors map[string]string

	// With make
	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	// Literal
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for c, v := range m {
		fmt.Println("Hes code for color", c, "is", v)
	}
}
