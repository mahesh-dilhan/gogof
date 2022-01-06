package main

type Chip interface {
	powerby(manfacture string)
}

type chip struct {
	manfacture string
}
