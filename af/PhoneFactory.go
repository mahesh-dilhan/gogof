package main

import "fmt"

type PhoneFactory interface {
	makeChip() Chip
}

func getPhoneFactory(model string) (PhoneFactory, error) {
	if model == "iphone" {
		return &iphone{}, nil
	}
	//if model == "nokia" {
	//	return &nokia{}, nil
	//}
	return nil, fmt.Errorf("Wrong model passed")
}
