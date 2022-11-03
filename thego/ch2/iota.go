package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	fmt.Printf("Wednesday=%d \t FlagLoopback=%d\t KiB=%d",
		Wednesday, FlagLoopback, KiB)
}

/*
Exercise 3.13: Write const declaration for KB, MB, up through YB as compactly as you can.
*/
