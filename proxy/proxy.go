package http_proxy

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var httpClint *http.Client

func init() {
	httpClint = &http.Client{}
}

type Transpond struct {
	Host           string
	TimeOut        time.Duration
	Retransmission int
	httpClient     *http.Client
}

func NewTranspond(host string, option ...Option) *Transpond {
	t := new(Transpond)
	t.Host = host

	if len(option) > 0 {
		for _, v := range option {
			v(t)
		}
	}
	return t
}

type Option func(t *Transpond)

func WithTomeOut(timeOut int) Option {
	return func(t *Transpond) {
		t.TimeOut = time.Duration(timeOut)
	}
}

func WithRetransmission(retransmission int) Option {
	return func(t *Transpond) {
		t.Retransmission = retransmission
	}
}

func (t *Transpond) TranspondGet(r *http.Request) (interface{}, error) {
	httpClint.Timeout = t.TimeOut
	urlStr := r.URL.String()
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, errors.New("New request err: " + err.Error())
	}
	req.Header = r.Header
	rep, err := httpClint.Do(req)
	if err != nil {
		return nil, errors.New("http Do err: " + err.Error())
	}
	return rep, nil
}

func (t *Transpond) TranspondPost(r *http.Request) (interface{}, error) {
	httpClint.Timeout = t.TimeOut
	urlStr := r.URL.String()
	req, err := http.NewRequest(http.MethodPost, urlStr, r.Body)
	if err != nil {
		return nil, errors.New("New request err: " + err.Error())
	}
	req.Header = r.Header
	rep, err := httpClint.Do(req)
	if err != nil {
		return nil, errors.New("http Do err: " + err.Error())
	}
	return rep, nil
}

func (t *Transpond) Proxy(w http.ResponseWriter, r *http.Request) error {
	u, err := url.Parse(t.Host)
	if err != nil {
		return errors.New("proxy parse err :" + err.Error())
	}
	rp := httputil.NewSingleHostReverseProxy(u)
	rp.ServeHTTP(w, r)
	return nil
}
