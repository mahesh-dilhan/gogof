package main

import "fmt"

func main() {
	iphoneFactory, _ := getPhoneFactory("iphone")
	iphoneChip := iphoneFactory.makeChip()
	printChipDetails(iphoneChip)

}

func printChipDetails(c Chip) {
	fmt.Printf("Power by: %s", c.powerby)
}
