package main

import (
	"encoding/json"
	"fmt"
)

var theFruit = map[string]string{
	"apple":  "🍎",
	"banana": "🍌",
	"cherry": "🍒",
}

type fruit struct {
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

func main() {
	var fruits []fruit

	for name, emoji := range theFruit {
		fruits = append(fruits, fruit{name, emoji})
	}

	fruitBytes, err := json.Marshal(fruits) // HL
	if err != nil {
		panic(err)
	}

	// type conversion
	fmt.Println(string(fruitBytes))

	var moreFruits []fruit
	err = json.Unmarshal(fruitBytes, &moreFruits) // HL

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", moreFruits)
}
