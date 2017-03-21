package lineshift

import (
	"os"
	"log"
	"bytes"
)

//noinspection GoUnusedExportedFunction
func Lineshift(filename string, search string) int{
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