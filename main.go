package main

import (
	"fmt"
)
func init() {
	subjectExample := observer.NewSubjectExample()
	observerExample := observer.NewObserverExample()

}
func main() {

	i, j := 0, 0
	fmt.Scanf("%d ", &i)
	fmt.Scanf("%d ", &j)
	done := make(chan int)

	go func() {
		done <- i + j
	}()

	fmt.Println(<-done)
	fmt.Println("here")

}
