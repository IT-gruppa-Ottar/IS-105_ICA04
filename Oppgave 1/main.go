package main

import (
	"fmt"
	"os"
	"./filutils"
	"./lineshift"
)

func main() {
	fmt.Println("Oppgave 1a")
	fmt.Println("Text 1")
	fileutils.FileToByteslice("text1.txt")

	fmt.Println("Text 2")
	fileutils.FileToByteslice("text2.txt")

	fmt.Println()

	fmt.Println("Oppgave 1b")

	filename := os.Args[1]
	amount := lineshift.Lineshift(filename, "\n")

	fmt.Println(amount)
}
