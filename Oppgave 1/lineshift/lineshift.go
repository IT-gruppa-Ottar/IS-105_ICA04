
//kode som vi kan bruke for å løse oppgaven

package lineshift

import (
	"fmt"
	"strings"

)


func LineShiftCheck()  {

	//Denne stringen skal egentlig returneres av en fil.
	// Får ikke til å bruke FileToByteSlice til å returnere de eksakte binære verdiene som en string, og kan derfor ikke bruke strings.contains da den ikke støtter binære verdier

	str := "84 101 115 116 101 114 32 108 105 110 106 101 115 107 105 102 116 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10"

	subStr := "10"

	if strings.Contains(str, subStr) {
		fmt.Println("Contains binary 10 = contains linefeed")
	} else {
		fmt.Println("Does not contain linefeed")
	}

	subStr2 := "13"

	if strings.Contains(str, subStr2) {
		fmt.Println("Contains binary 13 = contains carraige return")
	} else {
		fmt.Println("Does not contain carriage return")
	}

}

func LineShiftCheck2(filename string)  {

	//Fungerer ikke

	subStr := "\n+"

	if strings.Contains(filename, subStr) {
		fmt.Println("Contains linefeed")
	} else {
		fmt.Println("Does not contain linefeed")
	}

	subStr2 := "\r+"

	if strings.Contains(filename, subStr2) {
		fmt.Println("Contains carriage return")
	} else {
		fmt.Println("Does not contain carriage return")
	}

}