package model

type Gender string

type Student struct {
	ID         uint
	Name       string
	Email      string
	Age        uint
	City       City
	University University
	Father     Person
	Mother     Person
}

type Person struct {
	ID     uint
	Name   string
	Gender Gender
}
