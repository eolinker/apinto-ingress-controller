package setting

import (
	"context"
	"encoding/json"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/client"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/nettest"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type settings struct {
	data *v1.Setting
}

func (ro *settings) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if !strings.HasPrefix(r.URL.Path, "/api/setting/plugin") {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		resp := ro.get()
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(resp)
		_, _ = w.Write(data)
	case http.MethodPost:
		data, _ := ioutil.ReadAll(r.Body)
		var create v1.Setting
		err := json.Unmarshal(data, &create)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ro.update(&create)
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(create)
		_, _ = w.Write(res)
		return

	}
}

func (r *settings) get() *v1.Setting {
	return r.data
}

func (r *settings) update(setting *v1.Setting) {
	r.data = setting
}

func runSettingServer(t *testing.T) *http.Server {
	ln, _ := nettest.NewLocalListener("tcp")
	httpSrv := &http.Server{
		Addr: ln.Addr().String(),
		Handler: &settings{
			data: &v1.Setting{
				Name:       "Setting",
				Profession: "Setting",
				Driver:     "plugins",
				Plugins: v1.SettingPlugins{
					{
						ID:     "eolinker.com:apinto:extra_params",
						Name:   "extra_params",
						Type:   "service",
						Status: "global",
						Config: map[string]interface{}{
							"params": []struct {
								Name     string      `json:"name"`
								Position string      `json:"position"`
								Value    interface{} `json:"value"`
								Conflict string      `json:"conflict"`
							}{
								{
									Name:     "a",
									Position: "query",
									Value:    "1",
									Conflict: "Convert",
								},
							},
							"error_type": "text",
						},
					},
				},
			},
		},
	}
	go func() {
		if err := httpSrv.Serve(ln); err != nil && err != http.ErrServerClosed {
			t.Errorf("failed to run http server: %s", err)
		}
	}()
	return httpSrv
}
func TestSetting(t *testing.T) {
	ser := runSettingServer(t)
	defer func() {
		assert.Nil(t, ser.Shutdown(context.Background()))
	}()
	u := url.URL{
		Scheme: "http",
		Host:   ser.Addr,
		Path:   "/api",
	}
	cli, err := client.NewClient(u.String(), 0, "")
	if err != nil {
		t.Fatal(err)
	}
	r := NewSetting(cli)
	c := context.Background()
	// test get
	res, err := r.GetPlugin(c)
	assert.Nil(t, err)
	// test update
	res.Plugins = v1.SettingPlugins{
		{
			ID:     "eolinker.com:apinto:extra_params",
			Name:   "extra_params",
			Type:   "service",
			Status: "global",
		},
	}
	_, err = r.UpdatePlugin(c, res)
	assert.Nil(t, err)
	res2, err := r.GetPlugin(c)
	assert.Nil(t, err)
	assert.Equal(t, res, res2)
	t.Log("test update successfully")

}
