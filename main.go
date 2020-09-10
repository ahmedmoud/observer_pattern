package main

import (
	"fmt"
	"observer_pattern/Enums"
	"observer_pattern/observer"
	"time"
)

func main() {

	//subjectExample := observer.NewSubjectExample("new subject instance")
	//observerExample := observer.NewObserverExample()
	//subjectExample.RegisterObserver(observerExample)
	//subjectExample.NotifyObservers()
	observer.Listen(Enums.EXAMPLE_OBSERVER_KEY)
	channel := make(chan string, 2)

	 go func(channel chan string) {
		 //time.Sleep(2* time.Second)
		 fmt.Println("my channel number " + <-channel)
		 fmt.Println("my channel number " + <-channel)
		 fmt.Println("my channel number " + <-channel)

	 }(channel)
	fmt.Println("here 1")
	channel <- "1"
	fmt.Println("here 2")
	channel <- "2"
	channel <- "3"
	fmt.Println("here 3")

	time.Sleep(20 * time.Second)

}
