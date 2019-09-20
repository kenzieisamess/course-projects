package main

import "fmt"

func printMap(c map[string]string){
	for key, value := range c{
		fmt.Println(key + value)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")
	// fmt.Println(colors)

	printMap(colors)
}
