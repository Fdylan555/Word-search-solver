package main

import (
	"bufio"
	"fmt"

	"os"

	"github.com/01-edu/z01"
)

func main() {
	readFile, err :=
		os.Open("word.txt")
	if err != nil {
		fmt.Println(err)
		fileScanner :=
			bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			fmt.Print(fileScanner.Text())
		}
		readFile.Close()
	}
	var suj2 = [10][10]rune{
		{'c', 'o', 't', 'd', 't', 'r', 's', 'n', 'e', 'c'},
		{'r', 'e', 'e', 'o', 't', 'e', 'o', 'h', 'u', 'c'},
		{'ê', 'u', 'h', 'h', 't', 'r', 'l', 'a', 'o', 'a'},
		{'p', 'f', 'w', 'a', 'e', 'r', 'e', 'a', 'a', 'f'},
		{'e', 's', 't', 'p', 'r', 't', 'a', 'e', 'r', 'é'},
		{'s', 'a', 'p', 'y', 'u', 'i', 'o', 'e', 's', 'd'},
		{'p', 'a', 'e', 'o', 'i', 't', 'c', 'd', 't', 'e'},
		{'n', 'n', 'c', 'n', 'c', 'w', 'b', 'o', 'b', 'n'},
		{'f', 'o', 'u', 'r', 'c', 'h', 'e', 't', 't', 'e'},
		{'b', 'a', 'i', 'e', 's', 't', 'h', 'n', 'w', 's'},
	}
	z01.PrintRune('\n')
	fmt.Print(suj2)
	//tab1 := [10]rune{}
	//tab2 := [10]rune{}
	//for tab1; tab1 'a' <= 'z'; tab1++ {

	//}
}
