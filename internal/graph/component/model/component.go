package model

type Component struct {
	Name         string
	AbsoluteName string
	Dependencies []Dependency
}
