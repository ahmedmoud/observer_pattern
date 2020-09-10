package observer

import (
	"log"
)

var subject *SubjectExample

type SubjectExample struct {
	observerList   [] IObserver
	StringToUpdate string
}

func NewSubjectExample(newValue string) *SubjectExample {
	if subject != nil {
		return subject
	}
	subject = &SubjectExample{}
	subject.StringToUpdate = newValue
	return subject
}

func (subject *SubjectExample) RegisterObserver(observer IObserver) {
	key := observer.GetKey()
	if subject != nil {
		for _, item := range subject.observerList {
			if item.GetKey() == key {
				log.Fatal("error register same observer")
				return
			}
		}

		subject.observerList = append(subject.observerList, observer)
	}

}

func (subject *SubjectExample) NotifyObservers() {

	for _, item := range subject.observerList {
		if item.Async() {
			channel := make(chan string)
			go item.Update(channel)
		}
	}
	for _, item := range subject.observerList {
		if !item.Async() {
			channel := make(chan string)
			item.Update(channel)
		}
	}
}
