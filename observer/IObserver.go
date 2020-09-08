package observer


type IObserver interface {
	Update(subject ISubject)
	GetKey() string
	IsAsync() bool
}