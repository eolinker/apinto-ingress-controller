package ingress

type IIngressControllerEvent interface {
	// run while receive add event form kubernetes api server
	onAdd(obj interface{})
	// run while receive update event form kubernetes api server
	onUpdate(obj interface{})
	// run while receive delete event form kubernetes api server
	onDelete(obj interface{})
}
