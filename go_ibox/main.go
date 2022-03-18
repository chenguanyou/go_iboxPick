package main

import "fmt"

func main() {
	// 1、发送验证码
	//res, _ := SendSms("18888888888")
	//fmt.Printf(res.Message)
	///////////////////////////////

	//2、登录获取token
	//res, _ := Login("18888888888", "99999")
	//fmt.Println(res.Data["token"])

	var runIndex = 0
	for runIndex <= LowConfig.RunWork {
		go LowLook()
		fmt.Printf("线程%v启动 \n\n", runIndex)
		runIndex += 1
	}
	select {}
}
