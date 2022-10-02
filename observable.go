package awesomeProject1

type Observable interface {
	subscribe(observable Observable)
	unsubscribe(observable Observable)
	sendAll()
}
