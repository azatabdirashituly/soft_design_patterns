package main

import (
	"fmt"
)

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{name: name}
}
func (s *Subscriber) Update(youtuberName, videoName string) {
	fmt.Printf("%s received a notification: %s uploaded %s\n", s.name, youtuberName, videoName)
}
