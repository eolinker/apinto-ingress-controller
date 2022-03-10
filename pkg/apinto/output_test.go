package apinto

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/nettest"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
)

type outputs struct {
	mu   sync.RWMutex
	data map[string]*v1.Output
}

func (ro *outputs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if !strings.HasPrefix(r.URL.Path, "/api/output") {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		name := strings.TrimPrefix(r.URL.Path, "/api/output")
		if len(name) == 0 {
			// list
			resp := ro.list()
			w.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(resp)
			_, _ = w.Write(data)
		} else {
			name = strings.TrimPrefix(name, "/")
			resp, err := ro.get(name)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(resp)
			_, _ = w.Write(data)
		}

	case http.MethodPut:
		name := strings.TrimPrefix(r.URL.Path, "/api/output/")
		data, _ := ioutil.ReadAll(r.Body)
		var update v1.Output
		err := json.Unmarshal(data, &update)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		_ = ro.update(name, &update)
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(update)
		_, _ = w.Write(res)
		return
	case http.MethodPost:
		data, _ := ioutil.ReadAll(r.Body)
		var create v1.Output
		err := json.Unmarshal(data, &create)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		_ = ro.create(&create)
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(create)
		_, _ = w.Write(res)
		return
	case http.MethodDelete:
		name := strings.TrimPrefix(r.URL.Path, "/api/output/")
		d, err := ro.del(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(d)
		_, _ = w.Write(res)
		return

	}
}

func (ro *outputs) genID(name string) string {
	return fmt.Sprintf("%s@output", name)
}

func (r *outputs) get(name string) (*v1.Output, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.data[r.genID(name)]
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("output %s not exist", name)
}

func (r *outputs) list() []*v1.Output {
	r.mu.RLock()
	res := make([]*v1.Output, 0, len(r.data))
	for _, v := range r.data {
		res = append(res, v)
	}
	r.mu.RUnlock()
	return res
}
func (r *outputs) update(name string, output *v1.Output) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	output.ID = r.genID(name)
	r.data[output.ID] = output
	return output.ID
}
func (r *outputs) create(output *v1.Output) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	output.ID = r.genID(output.Name)
	r.data[output.ID] = output
	return output.ID
}
func (r *outputs) del(name string) (*v1.Output, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	v, ok := r.data[r.genID(name)]
	if !ok {
		return nil, fmt.Errorf("not exist")
	}
	delete(r.data, r.genID(name))
	return v, nil
}
func runOutputServer(t *testing.T) *http.Server {
	ln, _ := nettest.NewLocalListener("tcp")
	httpSrv := &http.Server{
		Addr: ln.Addr().String(),
		Handler: &outputs{
			data: make(map[string]*v1.Output),
		},
	}
	go func() {
		if err := httpSrv.Serve(ln); err != nil && err != http.ErrServerClosed {
			t.Errorf("failed to run http server: %s", err)
		}
	}()
	return httpSrv
}
func TestOutput(t *testing.T) {
	ser := runOutputServer(t)
	defer func() {
		assert.Nil(t, ser.Shutdown(context.Background()))
	}()
	u := url.URL{
		Scheme: "http",
		Host:   ser.Addr,
		Path:   "/api",
	}
	cli, err := NewClient(u.String(), 0, "")
	if err != nil {
		t.Fatal(err)
	}
	cases := []*v1.Output{
		{
			Metadata: v1.Metadata{
				Name:       "demo1",
				Profession: "output",
				Driver:     "nsqd",
			},
			Config: map[string]interface{}{
				"Topic": "test_nsqd",
			},
		},
		{
			Metadata: v1.Metadata{
				Name:       "demo2",
				Profession: "output",
				Driver:     "kafka",
			},
			Config: map[string]interface{}{
				"Topic": "test_kafka",
			},
		},
	}
	r := NewOutput(cli)
	c := context.Background()
	// test create
	id, err := r.Create(c, cases[0])
	assert.Nil(t, err)
	assert.Equal(t, "demo1@output", id)
	cases[0].ID = id
	id, err = r.Create(c, cases[1])
	assert.Nil(t, err)
	assert.Equal(t, "demo2@output", id)
	cases[1].ID = id
	t.Log("test create successfully")
	// test get
	res, err := r.Get(c, cases[0].Name)
	assert.Nil(t, err)
	assert.Equal(t, cases[0], res)
	t.Log("test get successfully")
	// test list
	list, err := r.List(c)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(list))
	t.Log("test list successfully")
	// test update
	cases[1].Config["partition_type"] = "robin"
	_, err = r.Update(c, cases[1])
	assert.Nil(t, err)
	res, err = r.Get(c, cases[1].Name)
	assert.Nil(t, err)
	assert.Equal(t, cases[1], res)
	t.Log("test update successfully")
	// test delete
	err = r.Delete(c, cases[1].Name)
	assert.Nil(t, err)
	err = r.Delete(c, cases[0].Name)
	assert.Nil(t, err)
	list, err = r.List(c)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))
	t.Log("test delete successfully")
}
