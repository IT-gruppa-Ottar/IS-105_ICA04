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
	fmt.Println("-------------------------------------------")
	fmt.Println("Information about file:", fileInfo.Name())

	if fileInfo.Size() < 1000 {
		fmt.Println("Size in bytes:", fileInfo.Size())

	} else if fileInfo.Size() >= 1000 && fileInfo.Size() <= 1000000{
		fmt.Println("Size in KB:", float32(fileInfo.Size()) / float32(1024))

	} else if fileInfo.Size() >= 1000000{
		fmt.Println("Size in MB:", float32(fileInfo.Size()) / float32(1048576))
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
	fmt.Println("-------------------------------------------")

}