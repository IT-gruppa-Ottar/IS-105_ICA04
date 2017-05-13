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
	kibi := size / 1024
	mebi := size / 1048576
	gibi := size / 1073741824
	fmt.Println("File size in Bytes")
	fmt.Printf("%.2f", size)
	fmt.Println()
	fmt.Println("File size in Kibibytes")
	fmt.Printf("%.2f", kibi)
	fmt.Println()
	fmt.Println("File size in Mebibytes")
	fmt.Printf("%.2f", mebi)
	fmt.Println()
	fmt.Println("File size in Gibibytes")
	fmt.Printf("%.2f", gibi)
	fmt.Println()



	fmt.Println()
	mode := fileInfo.Mode()
	fmt.Println("Is a directory: \t\t\t\t", mode&os.ModeDir != 0)
	fmt.Println("Is a regular file: \t\t\t\t", mode.IsRegular())
	fmt.Println("Has Unix permission bits: \t\t", fileInfo.Mode())
	fmt.Println("Is append only: \t\t\t\t", mode&os.ModeAppend != 0)
	fmt.Println("Is a device file: \t\t\t\t", mode&os.ModeDevice != 0)
	fmt.Println("Is a Unix character device: \t", mode&os.ModeCharDevice != 0)
	fmt.Println("Is a Unix block device: \t\t", mode&os.ModeSocket != 0)
	fmt.Println("Is a symbolic link: \t\t\t", mode&os.ModeSymlink != 0)

}