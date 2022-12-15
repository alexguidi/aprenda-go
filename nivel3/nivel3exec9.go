package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Print("futebol")
	case "voley":
		fmt.Print("voley")
	case "hurling":
		fmt.Print("hurling")
	}
}
