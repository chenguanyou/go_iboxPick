package main

import (
	"bytes"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

type sendSmsData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func SendSms(phone string) (sendSmsDataJson *sendSmsData, err error) {
	api := "https://www.ibox.art/nft-mall-web/nft/user/sendSMSCode"
	data := make(map[string]interface{})
	data["transId"] = uuid.NewV4()
	data["phoneNumber"] = phone
	if dataBytes, err := json.Marshal(data); err != nil {
		return sendSmsDataJson, err
	} else {
		dataReader := bytes.NewReader(dataBytes)
		if req, err := http.NewRequest("POST", api, dataReader); err != nil {
			return sendSmsDataJson, err
		} else {
			req.Header.Set("accept", Headers.accept)
			req.Header.Set("content-type", Headers.contentType)
			req.Header.Set("hb-nft-version", Headers.hbNftVersion)
			req.Header.Set("accept-encoding", Headers.acceptEncoding)
			req.Header.Set("hb-nft-os", Headers.hbNftOs)
			req.Header.Set("user-agent", Headers.userAgent)
			req.Header.Set("accept-language", Headers.acceptLanguage)
			client := http.Client{}
			if res, err := client.Do(req); err != nil {
				return sendSmsDataJson, err
			} else {
				if dataByte, err := ioutil.ReadAll(res.Body); err != nil {
					return sendSmsDataJson, err
				} else {
					sendSmsDataJson = &sendSmsData{}
					errJson := json.Unmarshal(dataByte, sendSmsDataJson)
					return sendSmsDataJson, errJson
				}
			}
		}
	}
}
