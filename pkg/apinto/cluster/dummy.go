package cluster

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession"
)

func NewDummy() *dummy {
	return &dummy{
		Dummy: profession.NewDummy(),
	}
}

type dummy struct {
	*profession.Dummy
}

func (d *dummy) RouterChecker() profession.Checker {
	return d
}

func (d *dummy) ServiceChecker() profession.Checker {
	return d
}

func (d *dummy) UpstreamChecker() profession.Checker {
	return d
}

func (d *dummy) DiscoveryChecker() profession.Checker {
	return d
}

func (d *dummy) OutputChecker() profession.Checker {
	return d
}

func (d *dummy) AuthChecker() profession.Checker {
	return d
}

func (d *dummy) SettingChecker() profession.Checker {
	return d
}
