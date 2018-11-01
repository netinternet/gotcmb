package main

import (
	"fmt"
)

func main() {
	data := make(map[string]string)
	data["osman"] = "arzu"
	data["makal"] = "rzyaeva"
	for i, v := range data {
		fmt.Println(data[v])
		fmt.Println(data[i])
	}
}
