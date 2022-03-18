package main

import (
	"bytes"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"time"
)

type loginData struct {
	Code    string                 `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
	MsgId   string                 `json:"msg_id"`
	Success bool                   `json:"success"`
}

func Login(phone, code string) (loginDataJson *loginData, err error) {
	api := "https://www.ibox.art/nft-mall-web/v1.1/nft/user/login"
	data := make(map[string]interface{})
	data["transId"] = uuid.NewV4()
	data["phoneNumber"] = phone
	data["code"] = code
	if dataBytes, err := json.Marshal(data); err != nil {
		return loginDataJson, err
	} else {
		dataReader := bytes.NewReader(dataBytes)
		if req, err := http.NewRequest("POST", api, dataReader); err != nil {
			return loginDataJson, err
		} else {
			req.Header.Set("accept", Headers.accept)
			req.Header.Set("content-type", Headers.contentType)
			req.Header.Set("hb-nft-version", Headers.hbNftVersion)
			req.Header.Set("accept-encoding", Headers.acceptEncoding)
			req.Header.Set("hb-nft-os", Headers.hbNftOs)
			req.Header.Set("user-agent", Headers.userAgent)
			req.Header.Set("accept-language", Headers.acceptLanguage)
			client := http.Client{
				Timeout: time.Second * 10,
			}
			if res, err := client.Do(req); err != nil {
				return loginDataJson, err
			} else {
				if dataByte, err := ioutil.ReadAll(res.Body); err != nil {
					return loginDataJson, err
				} else {
					loginDataJson = &loginData{}
					errJson := json.Unmarshal(dataByte, loginDataJson)
					return loginDataJson, errJson
				}
			}
		}
	}
}
