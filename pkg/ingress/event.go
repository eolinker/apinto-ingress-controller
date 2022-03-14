package ingress

type IIngressControllerEvent interface {
	// OnAdd Run while receive add event form kubernetes api server
	OnAdd(obj interface{})
	// OnUpdate Run while receive update event form kubernetes api server
	OnUpdate(obj interface{})
	// OnDelete Run while receive delete event form kubernetes api server
	OnDelete(obj interface{})
}
