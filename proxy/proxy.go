package proxy

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Proxy struct {
	
}

// request relayer http handler
func (p *Proxy) ProxyHandler(w http.ResponseWriter, r *http.Request) {
	// Log incoming request HEADER & BODY
	fmt.Println("Request Headers:", r.Header)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err.Error())
	} else {
		fmt.Println("Request Body:", string(body))
	}
	newBody := bytes.NewReader(body)


    // Forward the request from Consensus to Execution
	// TODO: start a geth at port 1234 to test
    req, _ := http.NewRequest(r.Method, "https://weathered-broken-pond.discover.quiknode.pro/92d7a1da0e2c55ffa42c3e677973b841c9d7f1b4", newBody)
    req.Header = r.Header
    client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Forward the response to Consensus
    for key, value := range resp.Header {
        w.Header().Set(key, value[0])
    }
    w.WriteHeader(resp.StatusCode)
	
	// Log responding/outgoing request HEADER & BODY
	fmt.Println("Response Headers:", resp.Header)
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err.Error())
	} else {
		fmt.Println("Response Body:", string(body))
	}
	newBody = bytes.NewReader(body)

	io.Copy(w, newBody)
}
