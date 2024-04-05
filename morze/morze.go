package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("morze.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var data string = "bl"

	for {
		n, err := file.Read(data)
		if err == io.EOF || n >= 0 { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(data, "\n")
	}
}
