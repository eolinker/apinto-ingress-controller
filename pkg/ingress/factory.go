package ingress

type manager struct {
}

func Register() {

}

func (m *manager) NewController(name string, workers int) IIngressController {
	//TODO implement me
	panic("implement me")
}
