package observer

import (
	"Enums"
	"fmt"
)

type Example struct {
	Key  string
	sync bool
}

func NewObserverExample() *Example {
	example := &Example{}
	example.Key = Enums.EXAMPLE_OBSERVER_KEY
	return example
}
func (example *Example) Update(subject ISubject) {

	subjectExample := subject.(*SubjectExample)
	fmt.Println(subjectExample.observerList)
}

func (example *Example) GetKey() string {
	return example.Key
}
func (example *Example) IsAsync() bool {
	return example.sync
}
