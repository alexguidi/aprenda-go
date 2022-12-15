package main

import "fmt"

func main() {
	maps := make(map[string][]string, 3)

	maps["guidi"] = []string{"Programming", "Music", "Games"}
	maps["Nyc"] = []string{"Dance", "Photo", "Series"}

	for k, v := range maps {

		fmt.Println(k, v)
	}
}
