package mycat

import (
	"bufio"
	"fmt"
	"os"
)

var lineNumber = 1
var isLineNumber bool
var isPrintFileName bool

func openFile(filename string) (*os.File, error) {
	sf, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return sf, nil
}

func printFile(scanFiles []*os.File, filePaths []string) {
	filePathsNum := 0
	for i := range scanFiles {
		if isPrintFileName {
			fmt.Printf("fileName: %s\n", filePaths[filePathsNum])
			filePathsNum++
		}
		scanner := bufio.NewScanner(scanFiles[i])
		for scanner.Scan() {
			if isLineNumber {
				fmt.Printf("%v: ", lineNumber)
				lineNumber++
			}
			fmt.Println(scanner.Text())
		}
	}
}

func showHelp() {
	fmt.Println("mycat is file showing command")
}
func MyCat(args []string) {
	var filePaths []string
	var sfs []*os.File

	// check args
	if len(args) == 1 {

	}
	for i := 1; i < len(args); i++ {
		if flag := string([]rune(args[i])[:1]); flag != "-" {
			filePaths = append(filePaths, args[i])
		} else {
			switch args[i] {
			case "-p":
				isPrintFileName = true
			case "-n":
				isLineNumber = true
			case "-h":
				showHelp()
				return
			default:
				fmt.Fprintf(os.Stderr, "%s is unknown flag\n", args[i])
			}
		}
	}

	for i := range filePaths {
		sf, err := openFile(filePaths[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s can not open !\n", filePaths[i])
		} else {
			sfs = append(sfs, sf)
		}
		defer sf.Close()
	}

	printFile(sfs, filePaths)
}
