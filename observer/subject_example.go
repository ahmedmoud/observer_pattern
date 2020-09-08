package observer

import (
	"log"
)

var subject *SubjectExample

type SubjectExample struct {
	observerList   [] IObserver
	params         map[string]interface{}
	StringToUpdate string
}

func NewSubjectExample() *SubjectExample {
	if subject != nil {
		return subject
	}
	subject = &SubjectExample{}
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
		if item.IsAsync() {
			go item.Update(subject)
		}
	}
	for _, item := range subject.observerList {
		if !item.IsAsync() {
			 item.Update(subject)
		}
	}
}
