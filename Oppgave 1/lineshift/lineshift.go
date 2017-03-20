
//kode som vi kan bruke for å løse oppgaven

package lineshift

import (
	"fmt"
	"strings"
	//"../filutils"
	//"bytes"
)

func LineShiftCheck()  {

	//Denne stringen skal være bytes som returneres av filen som leses. Må finne en måte å gjøre dette.
	str := "84 101 115 116 101 114 32 108 105 110 106 101 115 107 105 102 116 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10"

	subStr := "10"

	if strings.Contains(str, subStr) {
		fmt.Printf("Contains 10 = contains lineshift \n")
	} else {
		fmt.Printf("subStr is not in str \n")
	}

	subStr2 := "13"

	if strings.Contains(str, subStr2) {
		fmt.Printf("Contains 13 = containslineshift \n")
	} else {
		fmt.Printf("subStr2 is not in str \n")
	}

}


/*
//Testet med bytes istedet for string, men får feil
func LineShiftCheck2()  {


	bt := "84 101 115 116 101 114 32 108 105 110 106 101 115 107 105 102 116 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10 79 103 32 101 110 32 116 105 108 32 46 46 46 13 10"

	subBt := "10"

	if bytes.Contains(bt, subBt){
		fmt.Printf("Contains 10 = contains newline \n")
	} else {
		fmt.Printf("subStr is not in str \n")
	}


}
*/