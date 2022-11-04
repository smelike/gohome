package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March",
		4: "April", 5: "May", 6: "June", 7: "July",
		8: "August", 9: "September", 10: "October",
		11: "November", 12: "December"}

	Q2 := months[4:7]     // Q2's capacity is 12 - 4 + 1 = 9
	summer := months[6:9] // summer's capacity is 12 - 6 + 1 = 7
	fmt.Println(Q2, summer)

	endlessSummer := summer[:5] // equal to months[6:10]
	fmt.Println(endlessSummer)
	fmt.Println(Q2[0:9])
}
