package main

import (
	"bytes"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"time"
)

type createOrderData struct {
	Data    map[string]interface{} `json:"data"`
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
}

func CreateOrder(albumId, gId, price, gNum string) (createOrderDataJson *createOrderData, err error) {
	api := "https://www.ibox.art/nft-mall-web/v1/nft/order/create"
	data := make(map[string]interface{})
	data["albumId"] = albumId
	data["gId"] = gId
	data["transId"] = uuid.NewV4()
	data["price"] = price //价格
	data["gNum"] = gNum
	data["payChannel"] = 22 // 付款渠道
	data["type"] = 0        // 0是藏品，1是盲盒
	if dataBytes, err := json.Marshal(data); err != nil {
		return createOrderDataJson, err
	} else {
		dataReader := bytes.NewReader(dataBytes)
		if req, err := http.NewRequest("POST", api, dataReader); err != nil {
			return createOrderDataJson, err
		} else {
			req.Header.Set("accept", Headers.accept)
			req.Header.Set("content-type", Headers.contentType)
			req.Header.Set("hb-nft-version", Headers.hbNftVersion)
			req.Header.Set("hb-nft-token", Headers.hbNftToken)
			req.Header.Set("accept-encoding", Headers.acceptEncoding)
			req.Header.Set("hb-nft-os", Headers.hbNftOs)
			req.Header.Set("user-agent", Headers.userAgent)
			req.Header.Set("accept-language", Headers.acceptLanguage)
			client := http.Client{
				Timeout: time.Second * 10,
			}
			if res, err := client.Do(req); err != nil {
				return createOrderDataJson, err
			} else {
				if dataByte, err := ioutil.ReadAll(res.Body); err != nil {
					return createOrderDataJson, err
				} else {
					createOrderDataJson = &createOrderData{}
					errJson := json.Unmarshal(dataByte, createOrderDataJson)
					return createOrderDataJson, errJson
				}
			}
		}
	}
}
