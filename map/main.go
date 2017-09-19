package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// color:=make(map[string]string) //create an empty map
	// color["white"] = "#ffffff"
	// structName.white
	// delete(color, white) //delete items in the map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hEX FOR COLOR", color, "is,", hex)
	}
}
