package main

import "fmt"

func main() {
	iphoneFactory, _ := getPhoneFactory("iphone")
	iphoneChip := iphoneFactory.makeChip()
	printShoeDetails(iphoneChip)

}

func printChipDetails(c Chip) {
	fmt.Printf("Power by: %s", c.powerby)
}
