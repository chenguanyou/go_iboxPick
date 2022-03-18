package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type resultData struct {
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
	Msgid   string                 `json:"msgid"`
	Success bool                   `json:"success"`
}

func LowPrice() (resultDataJson *resultData, err error) {
	api := "https://www.ibox.art/nft-mall-web/nft/product/getResellList"
	data := make(map[string]interface{})
	data["pageSize"] = LowConfig.PageSize //每页显示多少，默认20
	data["transId"] = uuid.NewV4()        //请求标识
	data["origin"] = LowConfig.Origin     //0在售
	data["type"] = LowConfig.Type         // 0是藏品，1是盲盒
	data["page"] = LowConfig.Page         // 分页
	data["sort"] = LowConfig.Sort         //  0是最新，1是价格最低开始
	if dataBytes, err := json.Marshal(data); err != nil {
		return resultDataJson, err
	} else {
		dataReader := bytes.NewReader(dataBytes)
		if req, err := http.NewRequest("POST", api, dataReader); err != nil {
			return resultDataJson, err
		} else {
			req.Header.Set("accept", Headers.accept)
			req.Header.Set("content-type", Headers.contentType)
			req.Header.Set("hb-nft-version", Headers.hbNftVersion)
			req.Header.Set("accept-encoding", Headers.acceptEncoding)
			req.Header.Set("hb-nft-os", Headers.hbNftOs)
			req.Header.Set("user-agent", Headers.userAgent)
			req.Header.Set("accept-language", Headers.acceptLanguage)
			client := http.Client{}
			if Config.proxyIpType {
				getProxyIp, _ := ProxyIp()
				proxyIp := "http://" + strings.Split(getProxyIp.Result[0], ",")[0]
				fmt.Println("使用代理ip：", proxyIp)
				parseUri, _ := url.Parse(proxyIp)
				client = http.Client{
					Timeout: time.Second * 10,
					Transport: &http.Transport{
						Proxy:                 http.ProxyURL(parseUri),
						MaxIdleConnsPerHost:   10,
						ResponseHeaderTimeout: time.Second * time.Duration(5),
					},
				}
			}
			if res, err := client.Do(req); err != nil {
				//fmt.Println("失败")
				return resultDataJson, err
			} else {
				if dataByte, err := ioutil.ReadAll(res.Body); err != nil {
					//fmt.Println("失败")
					return resultDataJson, err
				} else {
					resultDataJson = &resultData{}
					errJson := json.Unmarshal(dataByte, resultDataJson)
					//fmt.Println("成功")
					return resultDataJson, errJson
				}
			}
		}
	}
}
