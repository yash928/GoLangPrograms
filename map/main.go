package main

import "fmt"

func main() {
	// mp:=make(map[string]string)
	mp := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
	}
	mp["white"] = "#ffffff"
	fmt.Println(mp)
}
