package engine

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"git.ana/dorbmon/ivpn-gui/log"
	"github.com/elazarl/goproxy"
)

func newClient(port int) *http.Client {
	u, err := url.Parse(fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		panic(err)
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) { return u, nil },
		},
	}
}
func httpProxyHandler(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	u, err := url.Parse(_defaultKey.Socks5Addr)
	if err != nil {
		panic(err)
	}
	p, err := strconv.Atoi(u.Port())
	if err != nil {
		panic(err)
	}
	c := newClient(p)
	resp, err := c.Do(req)
	if err != nil {
		log.Errorf("[HTTP PROXY]: %s", err.Error())
		return nil, nil
	}
	return req, resp
}
