package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Please provide a mode, string and a key")
	}
	mode := os.Args[1]
	key := os.Args[3]
	output := ""
	var input []byte
	var err error

	if mode == "e" {
		temp := os.Args[2]
		input = []byte(temp)
	} else if mode == "d" {
		input, err = hex.DecodeString(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Invalid mode: " + mode)
	}
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	fmt.Println(output)
}
