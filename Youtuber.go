package main

import (
	"fmt"
	"sync"
)

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
	wg := sync.WaitGroup{}
	for _, subs := range y.subscribers {
		wg.Add(1)
		subs.Update(y.youtuberName, videoName)
	}
	wg.Wait()
}
func (y *Youtuber) UploadVideo(videoName string) {
	fmt.Printf("%s uploaded a new video called: %s\n", y.youtuberName, videoName)
	y.notifyObservers(videoName)
}
