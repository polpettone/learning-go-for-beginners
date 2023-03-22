package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Tell me your name: ")
	name := readInput()
	fmt.Printf("Hallo %s", name)
}

func readInput() string {
	var reader = bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
