package core

import (
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"sync"
)

const (
	APFallBackEndpoint = "ap.spotify.com:443"
	APResolveEndpoint  = "http://apresolve.spotify.com/"
)

type APResolver struct {
	m      *sync.Mutex
	client *http.Client
	list   APList
}

type APList struct {
	APList []string `json:"ap_list"`
}

func NewAPResolver(dialer proxy.Dialer) *APResolver {
	if dialer == nil {
		dialer = proxy.Direct
	}
	cl := &http.Client{Transport: &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
			type resStruct struct {
				conn net.Conn
				err  error
			}
			type resChan chan resStruct
			var res = make(resChan)
			go func() {
				conn, err := dialer.Dial(network, addr)
				res <- resStruct{
					conn: conn,
					err:  err,
				}
			}()
			select {
			case <-ctx.Done():
				return nil, errors.New("dialer : canceled by client")
			case r := <-res:
				return r.conn, r.err
			}
		},
	}}
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
		if a.list.APList == nil {
			err = errors.New("APResolver : empty aplist")
			return
		}
		*resPtr = a.list.APList
		return nil
	})
	return
}

func (a *APResolver) Resolve() (err error) {
	tryEndpoint := func(endpoint string) (err error) {
		res, err := a.client.Get(endpoint)
		defer func() {
			if res != nil {
				if res.Body != nil {
					_ = res.Body.Close()
				}
			}
		}()
		if err != nil {
			return a.doJob(func() (err error) {
				return json.NewDecoder(res.Body).Decode(&a.list)
			})
		}
		return
	}
	if err = tryEndpoint(APResolveEndpoint); err != nil {
		if err = tryEndpoint(APFallBackEndpoint); err != nil {
			return errors.New("APResolver : resolve failed")
		} else {
			return nil
		}
	} else {
		return nil
	}
}
