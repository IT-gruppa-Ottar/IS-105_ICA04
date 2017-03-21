package filemode

import (
	"fmt"
	"os"
	"log"
)

var (
	fileInfo os.FileInfo
	err      error
)

//noinspection GoUnusedExportedFunction
func FileInfo(filename string) {
	// Stat returns file info. It will return
	// an error if there is no file.
	// fileInfo, err = os.Stat("text1.txt")
	fileInfo, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Information about file:", fileInfo.Name())

	size := float64(fileInfo.Size())
	kibi := size / 1000
	mebi := size / 1000000
	gibi := size / 1000000000
	fmt.Println("File size in Bytes: ", size)
	fmt.Println("File size in Kibibytes: ", kibi)
	if(mebi > 0.1) {
		fmt.Println("File size in Mebibytes: ", mebi)
	} else{
		fmt.Println("File is under 0,1 Mebibytes.")
	}
	if(gibi > 0.1) {
		fmt.Println("File size in Gibibytes: ", gibi)
	} else{
		fmt.Println("File is under 0,1 Gibibytes.")
	}


	fmt.Println()
	mode := fileInfo.Mode()
	fmt.Println("Is a directory:			", mode&os.ModeDir != 0)
	fmt.Println("Is a regular file:		", mode.IsRegular())
	fmt.Println("Has Unix permission bits:	", fileInfo.Mode())
	fmt.Println("Is append only:			", mode&os.ModeAppend != 0)
	fmt.Println("Is a device file:		", mode&os.ModeDevice != 0)
	fmt.Println("Is a Unix character device:	", mode&os.ModeCharDevice != 0)
	fmt.Println("Is a Unix block device:		", mode&os.ModeSocket != 0)
	fmt.Println("Is a symbolic link:		", mode&os.ModeSymlink != 0)

}