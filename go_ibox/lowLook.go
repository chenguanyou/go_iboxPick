package main

import (
	"fmt"
	"github.com/pkg/browser"
	"strconv"
	"strings"
	"time"
)

var GID = make(map[string]bool)

func LowLook() {
	for {
		data, err := LowPrice()
		//data, err := CreateOrder("100000393", "100638386", "34.7", "34779")
		if err != nil {
			fmt.Println("err", err)
			continue
		} else {
			fmt.Println("success", data.Success)
			if data.Success == false {
				continue
			}
			data_list := data.Data["list"].([]interface{})
			//fmt.Println(data_list)
			//fmt.Println(err)
			//if err {
			//	fmt.Println("异常，重试中...")
			//	continue
			//}
			for _, item := range data_list {
				item_val := item.(map[string]interface{})
				gName := item_val["gName"].(string)
				albumName := item_val["albumName"].(string)
				albumId := strconv.Itoa(int(item_val["albumId"].(float64)))
				gId := item_val["gId"].(string)
				priceCny := item_val["priceCny"].(string)
				priceCny1 := strings.Split(priceCny, ".")[0]
				gNum := item_val["gNum"].(string)
				//albumIds := strconv.FormatFloat(albumId, 'E', -1, 64)
				fmt.Println(albumId)
				fmt.Println(gNum)
				fmt.Printf("gName:%v,albumName:%v,albumId:%v,gId:%v,priceCny:%v,gNum:%s\n\n", gName, albumName, albumId, gId, priceCny, gNum)
				prices, _ := strconv.ParseInt(priceCny1, 10, 64)
				fmt.Printf("%v, %v, %v <= %v, %v \n\n", prices, LowConfig.LowPrice, prices, LowConfig.LowPrice, prices <= LowConfig.LowPrice)
				if GID[gId] {
					fmt.Println("已经购买过了")
					continue
				}
				if prices <= LowConfig.LowPrice {
					GID[gId] = true
					cdata, cerr := CreateOrder(albumId, gId, priceCny, gNum)
					if cerr != nil {
						fmt.Println("提交订单发生了异常", cerr)
						break
					} else {
						if cdata.Success == false {
							fmt.Println("cdata.Data 提交订单发送异常", cdata)
							break
						}
						fmt.Println("cdata.Data", cdata.Data)
						fmt.Println("cdata.Data[orderStr]", cdata.Data["orderStr"])
						browser.OpenURL(cdata.Data["orderStr"].(string))
						if LowConfig.VideoType {
							VideoMp3()
						}
						break
					}
				}
			}
		}
		if LowConfig.Sletime > 0 {
			fmt.Printf("延时：%v 秒\n", LowConfig.Sletime)
			time.Sleep(time.Duration(LowConfig.Sletime) * time.Second)
		}
	}
}
