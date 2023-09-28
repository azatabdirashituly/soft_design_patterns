package main

type Observer interface {
	Update(name string, msg string)
}
