# go_iboxPick v1.0.0
go语言版本的ibox的捡漏工具，go环境版本 v1.16


# 项目目录说明
├── config.go 配置文件  
├── create_order.go 实现提交订单功能  
├── login.go 实现ibox的登录功能  
├── lowLook.go 捡漏的主功能入口  
├── low_price.go 实现价格从小到大排列  
├── main.go 入口文件  
├── ok.mp3 音频，用来实现捡漏成功后播放音乐  
├── proxyip.go 代理ip功能，可以在config里面配置接口  
├── send_sms.go 实现ibox发送短信验证码功能  
└── video_mp3.go 实现捡漏成功后播放音频  

# 运行
```
1、安装go环境
2、编辑config文件，进行设置配置信息。通过send_sms.go发送验证码，通过login.go进行登录获取到token设置到config.go中的HbNftToken项
3、运行go run main.go 
ps：第三步如果运行失败可以使用(go run main.go config.go create_order.go login.go low_price.go lowLook.go proxyip.go send_sms.go video_mp3.go)
```
# 编译
```
1、在go环境安装完成后、配置文件编辑好后的情况下使用go build main.go 进行编译
ps: 如果运行失败可以使用(go build main.go config.go create_order.go login.go low_price.go lowLook.go proxyip.go send_sms.go video_mp3.go)
```

# 项目会长期更新（觉得不错的可以打赏一下，谢谢支持）

<img src="./zfb.jpeg" width="25%" />

<img src="./wx.jpeg" width="25%" />
