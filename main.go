package main

import "fmt"

func main () {
	colors := map[string]string{
		"red": "#ff0000",
		"white": "#fff",
		"black": "#000",
	}

	printMap(colors)
}


func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println("The hex code of", color, "is", hex)
	}
}