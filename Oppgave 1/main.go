package main

import (
	"./filutils"
	"./lineshift"
	"fmt"
	"os"
)



func main() {
	fmt.Println("Oppgave 1a")

	fmt.Println("Text 1")
	fileutils.FileToByteslice("text1.txt")

	fmt.Println("Text 2")
	fileutils.FileToByteslice("text2.txt")



	fmt.Println("Oppgave 1b metode 1")
	lineshift.LineShiftCheck()

	fmt.Println("Oppgave 1b metode 2, denne metoden virker ikke, men tar fil som argument")
	filename := os.Args[1]
	lineshift.LineShiftCheck2(filename)




}
