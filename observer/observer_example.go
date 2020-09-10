package observer

import (
	"observer_pattern/Enums"
	"time"
)

type Example struct {
	Key  string
	sync bool
	subject SubjectExample
}

func NewObserverExample() *Example {
	example := &Example{}
	example.Key = Enums.EXAMPLE_OBSERVER_KEY
	example.sync = true
	return example
}
func (example *Example) Update(c chan string)  {
	time.Sleep(1 * time.Second)
	c <- example.Key
}

func (example *Example) GetKey() string {
	return example.Key
}
func (example *Example) Async() bool {
	return example.sync
}
