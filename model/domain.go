package model

// //Wg 线程同步函数
// var Wg sync.WaitGroup

//ExitCh 退出管道
var ExitCh = make(chan bool, 3)

//ResultCh Doamin结构体结果输出管道
var ResultCh = make(chan *Domain, 2100)

//Domain 域名结构体
type Domain struct {
	OrderID    string
	Name       string
	Suffix     string
	IsRegister bool
	Addr       string
	Balance    string
	UserAddr   string
}

//ResDomain 接口返回结构体
type ResDomain struct {
	Avail int    `json:"avail"`
	Name  string `json:"name"`
	Tld   string `json:"tld"`
}

//AliyunDomain 接口返回结构体
type AliyunDomain struct {
	Success   bool        `json:"success"`
	Module    []ResDomain `json:"module"`
	ErrorCode int         `json:"errorCode"`
}
