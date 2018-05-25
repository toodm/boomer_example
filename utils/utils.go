package utils

import (
	"bytes"
	"crypto/tls"
	//	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var Client *http.Client

//var timeout int

//var disableCompression bool
//var disableKeepalive bool

func init() {
	//	flag.IntVar(&timeout, "timeout", 10, "Seconds to max. wait for each response")
	//	flag.BoolVar(&disableCompression, "disable-compression", false, "Disable compression")
	//	flag.BoolVar(&disableKeepalive, "disable-keepalive", true, "Disable keepalive")
	//	flag.Parse()

	//	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2000
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConnsPerHost: 2000,
		DisableCompression:  false,
		DisableKeepAlives:   false,
	}
	Client = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(10) * time.Second,
	}
}

func HttpRequest(mothod string, url string, urlParam map[string]string, body []byte) (respBody []byte, err error) {
	var (
		req  *http.Request
		resp *http.Response
		//		respBody []byte
		//		err      error
	)
	req, err = http.NewRequest(mothod, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	//设置url参数
	vars := req.URL.Query()
	for url_k, url_v := range urlParam {
		vars.Add(url_k, url_v)
	}
	req.URL.RawQuery = vars.Encode()
	fmt.Println(req.URL.RawQuery)
	// Fetch Request
	resp, err = Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Read Response Body
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
