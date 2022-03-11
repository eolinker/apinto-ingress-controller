package profession

import (
	"context"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

func NewDummy() *Dummy {
	return &Dummy{
		routerDummy:    &routerDummy{},
		upstreamDummy:  &upstreamDummy{},
		serviceDummy:   &serviceDummy{},
		outputDummy:    &outputDummy{},
		discoveryDummy: &discoveryDummy{},
		authDummy:      &authDummy{},
		settingDummy:   &settingDummy{},
	}
}

type Dummy struct {
	routerDummy    *routerDummy
	serviceDummy   *serviceDummy
	settingDummy   *settingDummy
	upstreamDummy  *upstreamDummy
	discoveryDummy *discoveryDummy
	authDummy      *authDummy
	outputDummy    *outputDummy
}

func (d *Dummy) Router() Router {
	return d.routerDummy
}

func (d *Dummy) Upstream() Upstream {
	return d.upstreamDummy
}

func (d *Dummy) Service() Service {
	return d.serviceDummy
}

func (d *Dummy) Discovery() Discovery {
	return d.discoveryDummy
}

func (d *Dummy) Output() Output {
	return d.outputDummy
}

func (d *Dummy) Auth() Auth {
	return d.authDummy
}

func (d *Dummy) Setting() Setting {
	return d.settingDummy
}

func (d *Dummy) DelCheck(name string) (*response.Response, error) {
	return nil, ErrClusterNotExist
}

func (d *Dummy) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	return nil, ErrClusterNotExist
}

type routerDummy struct {
}

func (r *routerDummy) List(ctx context.Context) ([]*response.Response, error) {
	return nil, ErrClusterNotExist
}

func (r *routerDummy) Get(ctx context.Context, name string) (*v1.Router, error) {
	return nil, ErrClusterNotExist
}

func (r *routerDummy) Delete(ctx context.Context, name string) error {
	return ErrClusterNotExist
}

func (r *routerDummy) Update(ctx context.Context, router *v1.Router) (string, error) {
	return "", ErrClusterNotExist
}

func (r *routerDummy) Create(ctx context.Context, router *v1.Router) (string, error) {
	return "", ErrClusterNotExist
}

type serviceDummy struct {
}

func (s *serviceDummy) List(ctx context.Context) ([]*response.Response, error) {
	return nil, ErrClusterNotExist
}

func (s *serviceDummy) Get(ctx context.Context, name string) (*v1.Service, error) {
	return nil, ErrClusterNotExist
}

func (s *serviceDummy) Delete(ctx context.Context, name string) error {
	return ErrClusterNotExist
}

func (s *serviceDummy) Update(ctx context.Context, service *v1.Service) (string, error) {
	return "", ErrClusterNotExist
}

func (s *serviceDummy) Create(ctx context.Context, service *v1.Service) (string, error) {
	return "", ErrClusterNotExist
}

type upstreamDummy struct {
}

func (u *upstreamDummy) List(ctx context.Context) ([]*response.Response, error) {
	return nil, ErrClusterNotExist
}

func (u *upstreamDummy) Get(ctx context.Context, name string) (*v1.Upstream, error) {
	return nil, ErrClusterNotExist
}

func (u *upstreamDummy) Delete(ctx context.Context, name string) error {
	return ErrClusterNotExist
}

func (u *upstreamDummy) Update(ctx context.Context, upstream *v1.Upstream) (string, error) {
	return "", ErrClusterNotExist
}

func (u *upstreamDummy) Create(ctx context.Context, upstream *v1.Upstream) (string, error) {
	return "", ErrClusterNotExist
}

type outputDummy struct {
}

func (o *outputDummy) List(ctx context.Context) ([]*response.Response, error) {
	return nil, ErrClusterNotExist
}

func (o *outputDummy) Get(ctx context.Context, name string) (*v1.Output, error) {
	return nil, ErrClusterNotExist
}

func (o *outputDummy) Delete(ctx context.Context, name string) error {
	return ErrClusterNotExist
}

func (o *outputDummy) Update(ctx context.Context, output *v1.Output) (string, error) {
	return "", ErrClusterNotExist
}

func (o *outputDummy) Create(ctx context.Context, output *v1.Output) (string, error) {
	return "", ErrClusterNotExist
}

type settingDummy struct {
}

func (s *settingDummy) GetPlugin(ctx context.Context) (*v1.Setting, error) {
	return nil, ErrClusterNotExist
}

func (s *settingDummy) UpdatePlugin(ctx context.Context, setting *v1.Setting) (string, error) {
	return "", ErrClusterNotExist
}

type discoveryDummy struct {
}

func (d *discoveryDummy) List(ctx context.Context) ([]*response.Response, error) {
	return nil, ErrClusterNotExist
}

func (d *discoveryDummy) Get(ctx context.Context, name string) (*v1.Discovery, error) {
	return nil, ErrClusterNotExist
}

func (d *discoveryDummy) Delete(ctx context.Context, name string) error {
	return ErrClusterNotExist
}

func (d *discoveryDummy) Update(ctx context.Context, discovery *v1.Discovery) (string, error) {
	return "", ErrClusterNotExist
}

func (d *discoveryDummy) Create(ctx context.Context, discovery *v1.Discovery) (string, error) {
	return "", ErrClusterNotExist
}

type authDummy struct {
}

func (d *authDummy) List(ctx context.Context) ([]*response.Response, error) {
	return nil, ErrClusterNotExist
}

func (d *authDummy) Get(ctx context.Context, name string) (*v1.Auth, error) {
	return nil, ErrClusterNotExist
}

func (d *authDummy) Delete(ctx context.Context, name string) error {
	return ErrClusterNotExist
}

func (d *authDummy) Update(ctx context.Context, auth *v1.Auth) (string, error) {
	return "", ErrClusterNotExist
}

func (d *authDummy) Create(ctx context.Context, auth *v1.Auth) (string, error) {
	return "", ErrClusterNotExist
}
