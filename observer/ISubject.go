package observer


type ISubject interface {
	NotifyObservers()
	RegisterObserver(IObserver)
}
