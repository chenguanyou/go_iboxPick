package main

// 代理ip

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type proxyIp struct {
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
	Code      int      `json:"code"`
	Result    []string `json:"result"`
	Timestamp int      `json:"timestamp"`
}

func ProxyIp() (proxyIpJson *proxyIp, err error) {
	api := LowConfig.ProxyIp
	if req, err := http.Get(api); err != nil {
		return proxyIpJson, err
	} else {
		var proxyIpByte []byte
		if proxyIpByte, err = ioutil.ReadAll(req.Body); err != nil {
			return proxyIpJson, err
		} else {
			proxyIpJson = &proxyIp{}
			errJson := json.Unmarshal(proxyIpByte, proxyIpJson)
			return proxyIpJson, errJson
		}
	}
}
