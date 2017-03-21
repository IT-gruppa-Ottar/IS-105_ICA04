package main

import (
	"os"
	"fmt"
	"log"
	"bytes"
	"sort"
)

func main() {
	filename := os.Args[1]

	//Teller antall linjer i filen
	amount := searchFile(filename, "\n")
	fmt.Println("Totalt antall linjer:")
	fmt.Println(amount)

	//Finner de 5 mest brukte runene og viser hvor mange ganger de blir brukt
	temp := runes(filename)
	sortRunes(temp)
}

func searchFile(filename string, search string) int{
	//Åpner filen
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 512*1024)
	//Teller linjeskift
	counter := 0
	//Søker på linjeskift
	searchFor := []byte(search)

	c, _ := file.Read(buffer)
	counter += bytes.Count(buffer[:c], searchFor)
	return counter
}

func runes(filename string) map[int]string{
	//Lager map
	m := make(map[int]string)
	counter := 0

	//Runer for Ascii
	for i := 0x20; i <= 0x7F; i++ {
		counter = searchFile(filename, string(i))
		rune := string(i)
		m[counter] = rune
	}
	//fmt.Println(counter)
	return m
}

func sortRunes(m map[int]string){
	var keys []int

	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println()
	fmt.Println("5 runer som forekommer mest:")

	v1 := keys[len(keys)-1]
	v2 := keys[len(keys)-2]
	v3 := keys[len(keys)-3]
	v4 := keys[len(keys)-4]
	v5 := keys[len(keys)-5]

	fmt.Println("Antall:", v1, "Rune:", m[v1])
	fmt.Println("Antall:", v2, "Rune:", m[v2])
	fmt.Println("Antall:", v3, "Rune:", m[v3])
	fmt.Println("Antall:", v4, "Rune:", m[v4])
	fmt.Println("Antall:", v5, "Rune:", m[v5])
}
