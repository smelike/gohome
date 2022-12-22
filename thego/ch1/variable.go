package main

var a = "G" // package level variable, but not exported variable

func main() {
	n()
	m()
	n()
}

func n() { print(a) }

func m() {
	// a := "O" //
	a = "O" // it means different, its value is replaced by "O"
	print(a)
}
