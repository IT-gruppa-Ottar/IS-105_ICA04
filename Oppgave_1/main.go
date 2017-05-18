package main

import (
	"./filutils"
	"./lineshift"
	"fmt"
	"os"
)

/**
Main importerer funksjoner fra pakkene "filutils" og "lineshift" og kaller dem.

Oppgave 1a tar de 2 "txt" filene og kjører dem igjennom "filutils-go" Da får vi
en byteslice for hver av dem og kan sammenligne.

Oppgave 1b krever et filnavn som paramenter. "lineshift.go" sjekker filen for hvilken
type linjeskift filen har og returneres som true eller false.
 */
func main() {
	fmt.Println("-----Oppgave 1a-----")

	fmt.Println("Text 1")
	fmt.Println(fileutils.FileToByteslice("text1.txt"))

	fmt.Println("Text 2")
	fmt.Println(fileutils.FileToByteslice("text2.txt"))

	fmt.Println()

	fmt.Println("-----Oppgave 1b-----")
	filename := os.Args[1]
	lineshift.LineShift(filename)
}