package main

func main() {
	youtuber1 := NewYoutuber("ItProger")
	subs1 := NewSubscriber("user123")
	youtuber1.subscribe(subs1)
	youtuber1.UploadVideo("Изучение GoLang")
}
