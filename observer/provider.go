package observer

import (
	"observer_pattern/Enums"
)

type fn func() ISubject

var eventListener = map[string]fn{

	Enums.EXAMPLE_OBSERVER_KEY: func() ISubject {
		subjectExample := NewSubjectExample("new instance")
		observerExample := NewObserverExample()
		subjectExample.RegisterObserver(observerExample)
		return subject
	},
}

func Listen(eventKey string) {
	subject := eventListener[eventKey]()
	 subject.NotifyObservers()
}
