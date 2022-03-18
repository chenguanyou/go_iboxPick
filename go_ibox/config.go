package main

type config struct {
	version     string `json:"version"`
	userAgent   string `json:"user_agent"`
	hbNftOs     string `json:"hb_nft_os"`
	HbNftToken  string `json:"hb_nft_token"`
	proxyIpType bool   `json:"proxy_ip_type"`
}

type headers struct {
	accept         string `json:"accept"`
	contentType    string `json:"content_type"`
	hbNftVersion   string `json:"hb_nft_version"`
	hbNftToken     string `json:"hb_nft_token"`
	acceptEncoding string `json:"accept_encoding"`
	hbNftOs        string `json:"hb_nft_os"`
	userAgent      string `json:"user_agent"`
	acceptLanguage string `json:"accept_language"`
}

type lowConfig struct {
	ProxyIp   string `json:"proxy_ip"`
	RunWork   int    `json:"run_work"`
	Sletime   int    `json:"sletime"`
	VideoType bool   `json:"video_type"`
	LowPrice  int64  `json:"low_price"`
	PageSize  int    `json:"page_size"`
	Origin    int    `json:"origin"`
	Type      int    `json:"type"`
	Page      int    `json:"page"`
	Sort      int    `json:"sort"`
}

var LowConfig = &lowConfig{
	ProxyIp:   "",    //代理ip接口，目前只支持星速云的(xingsudaili.com)
	VideoType: false, // 抢到后是否自动播放音乐
	LowPrice:  100,   // 低于/等于多少钱后进行捡漏
	PageSize:  20,    // 每页显示多少，显示的越小，速度越快。
	RunWork:   1,     // 默认开始线程的数量，在不开启代理ip的情况下，超过一个会被屏蔽ip
	Sletime:   0,     // 延时的时间，开启代理ip的情况下，不需要设置延迟。(延迟单位是秒)
	Origin:    0,     // 在售，不要更改
	Type:      0,     // 0是藏品，1是盲盒，不要更改。
	Page:      1,     // 分页，第几页
	Sort:      1,     //  0是最新，1是价格最低开始
}

var Config = &config{
	version:     "1.0.6",                                     // app版本
	userAgent:   "iBox/1.0.0 (iPhone; iOS 15.4; Scale/3.00)", // 要模拟的型号
	hbNftOs:     "ios",                                       // 当前手机型号
	HbNftToken:  "",                                          // 账户的Toekn
	proxyIpType: true,                                        // 是否开启代理IP，代理ip需要自己去购买(代理ip仅支持刷新的时候使用，下单的时候默认使用本地ip)
}

var Headers = &headers{
	accept:         "*/*",
	contentType:    "application/json",
	hbNftVersion:   Config.version,
	hbNftToken:     Config.HbNftToken,
	acceptEncoding: "gzip, deflate, br",
	hbNftOs:        Config.hbNftOs,
	userAgent:      Config.userAgent,
	acceptLanguage: "zh-Hans-CN;q=1",
}
