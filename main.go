package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	length := os.Args[1]
	n, err := strconv.Atoi(length)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(PasswordGenerator(n))
}
