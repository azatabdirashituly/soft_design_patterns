package main

type Subscriber struct {
	name string
}

type Subject interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	notifyObservers(msg string)
}
