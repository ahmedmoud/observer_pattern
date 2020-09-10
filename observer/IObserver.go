package observer


type IObserver interface {
	Update(chan string)
	GetKey() string
	Async() bool
}