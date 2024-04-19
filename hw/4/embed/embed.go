package main

import (
	"embed"
)

//go:embed files/single_file.txt
var fileString string

//go:embed files/single_file.txt
var fileByte []byte

//go:embed files/*.hash
var folder embed.FS

func main() {
	println(fileString)
	println(string(fileByte))
	content1, _ :=
		folder.ReadFile("files/file1.hash")
	println(string(content1))
	content2, _ :=
		folder.ReadFile("files/file2.hash")
	println(string(content2))
}
