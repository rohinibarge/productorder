package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/rohinibarge/productorder/api"
)

var DefaultDialer = &net.Dialer{}

type trasport struct {
	roundTrip http.RoundTripper
	dialer    *net.Dialer
	initCon   time.Time
	endCon    time.Time
	req       time.Time
	endReq    time.Time
}

func GetHttpClient() *http.Client {
	tp := custTransport()
	client := &http.Client{Transport: tp}
	return client
}

func custTransport() *trasport {

	tr := &trasport{
		dialer: &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	}
	tr.roundTrip = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                tr.dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	return tr
}

func (tr *trasport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.req = time.Now()
	resp, err := tr.roundTrip.RoundTrip(r)
	if err != nil {
		fmt.Println("error occured while roundtrip", err)
	}
	tr.endReq = time.Now()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error occured while roundtrip read", err)
	}
	p := api.ProductList{}
	if err := json.Unmarshal(body, &p); err != nil { // Parse []byte to the go struct pointer
		fmt.Printf("Can not unmarshal JSON, err: %v", err)
	}
	for i := range p.Products {
		j := rand.Intn(i + 1)
		p.Products[i], p.Products[j] = p.Products[j], p.Products[i]
	}
	b, _ := json.Marshal(p)
	resp.Body = io.NopCloser(bytes.NewBuffer(b))
	return resp, err
}

func (tr *trasport) dial(network, addr string) (net.Conn, error) {
	tr.initCon = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.endCon = time.Now()
	return cn, err
}

func (tr *trasport) ReqDuration() time.Duration {
	return tr.Duration() - tr.ConnDuration()
}

func (tr *trasport) ConnDuration() time.Duration {
	return tr.endCon.Sub(tr.initCon)
}

func (tr *trasport) Duration() time.Duration {
	return tr.endReq.Sub(tr.req)
}
