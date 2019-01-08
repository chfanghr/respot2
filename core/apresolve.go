package core

import (
	"errors"
	"golang.org/x/net/proxy"
	"net/http"
	"sync"
)

const (
	APFallBack        = "ap.spotify.com:443"
	APResolveEndpoint = "http://apresolve.spotify.com/"
)

type APResolver struct {
	m      *sync.Mutex
	client *http.Client
	APList []string `json:"ap_list"`
}

func NewAPResolver(dialer proxy.Dialer) *APResolver {
	if dialer == nil {
		dialer = proxy.Direct
	}
	cl := &http.Client{Transport: &http.Transport{Dial: dialer.Dial}}
	return &APResolver{client: cl, m: new(sync.Mutex)}
}

func (a *APResolver) doJob(job func() error) error {
	a.m.Lock()
	defer a.m.Unlock()
	return job()
}

func (a *APResolver) GetAPList() (res []string, err error) {
	resPtr := &res
	err = a.doJob(func() (err error) {
		if res == nil {
			err = errors.New("APResolver : empty aplist")
			return
		}
		*resPtr = a.APList
		return nil
	})
	return
}
