//go:generate go run sum.go

package main

func main() {
	println("Ky, world!")
}

func init() {
	a := 1
	b := 4
	println(a + b)
}
