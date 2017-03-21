package main

import (
	"./filutils"
	"./lineshift"
	"fmt"
	"os"
)



func main() {
	fmt.Println("-----Oppgave 1a-----")

	fmt.Println("Text 1")
	fmt.Println(fileutils.FileToByteslice("text01.txt"))

	fmt.Println("Text 2")
	fmt.Println(fileutils.FileToByteslice("text02.txt"))

	fmt.Println()


	fmt.Println("-----Oppgave 1b-----")
	filename := os.Args[1]
	lineshift.LineShift(filename)
}
