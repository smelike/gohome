package main

import "fmt"

/* named return values

func func_name(param_name type ...) int {
	return integernumber
}
func func_name(param_name type ...) (string, string) {
	return "string1", "string2"
}
func func_name(param_name type ...) (x int, y int) {
	x = 100
	y = 200
	return
}
func func_name(param_name type ...) (x, y int) {
	x = 100
	y = 200
	return
}

*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	/*
		A return statement without arguments returns the named return values. This is known as a "naked" return.

		Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
	*/
	return
}

func g_main() {
	fmt.Println(split(27))
}
