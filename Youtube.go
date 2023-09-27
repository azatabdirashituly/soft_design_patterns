package main

import "fmt"

type Subject interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	notifyObservers(msg string)
}

type Observer interface {
	Update(name string, msg string)
}

type Youtuber struct {
	youtuberName string
	subscribers  []Observer
}

func NewYoutuber(name string) *Youtuber {
	return &Youtuber{
		youtuberName: name,
	}
}
func (y *Youtuber) subscribe(observer Observer) {
	y.subscribers = append(y.subscribers, observer)
}
func (y *Youtuber) unsubscribe(observer Observer) {
	for i, subs := range y.subscribers {
		if subs == observer {
			y.subscribers = append(y.subscribers[:i], y.subscribers[i+1:]...)
			break
		}
	}
}
func (y *Youtuber) notifyObservers(videoName string) {
	for _, subs := range y.subscribers {
		subs.Update(y.youtuberName, videoName)
	}
}
func (y *Youtuber) UploadVideo(videoName string) {
	fmt.Printf("%s uploaded a new video called: %s\n", y.youtuberName, videoName)
	y.notifyObservers(videoName)
}

type Subscriber struct {
	name string
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{name: name}
}
func (s *Subscriber) Update(youtuberName, videoName string) {
	fmt.Printf("%s received a notification: %s uploaded %s\n", s.name, youtuberName, videoName)
}

func main() {
	youtuber1 := NewYoutuber("ItProger")
	subs1 := NewSubscriber("user123")
	youtuber1.subscribe(subs1)
	youtuber1.UploadVideo("Изучение GoLang")
}
