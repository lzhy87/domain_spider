package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lzhy87/domain_spider/model"
	"github.com/lzhy87/domain_spider/utils"
)

//createString 获取2个字母的组合+后缀（.vip、.top、.io）。并放入*domain结构体,返回一个[]*domain结构体的切片
func createString() []*model.Domain {
	//https://www.west.cn/main/whois.asp?act=query&domains=azex&suffixs=.top&du=&v=0.7122194239976691
	//https://wanwang.aliyun.com/domain/searchresult/#/?keyword=aaex&suffix=vip
	//https://checkapi.aliyun.com/check/checkdomain?domain=zaex.vip&token=Y4119c3d7117497d680fae8277165a014
	//{"success":"true","module":[{"avail":1,"name":"zaex.vip","tld":"vip"}],"errorCode":0}
	/*AliyunDomain{
		Success bool `json:"success"`
		Module []ResDomain `json:"module"`
		ErrorCode int `json:"errorCode"`
	}*/
	var ds = make([]*model.Domain, 0)
	var sum int
	for i := 1; i <= 3; i++ {
		switch i {
		case 1:
			for i := 'a'; i <= 'z'; i++ {
				for j := 'a'; j <= 'z'; j++ {
					sum++
					domain := new(model.Domain)
					num := fmt.Sprintf("%c%c%s", i, j, "ex")
					domain.Name = num
					domain.Suffix = "vip"
					domain.OrderID = strconv.Itoa(sum)
					url := fmt.Sprintf("https://checkapi.aliyun.com/check/checkdomain?domain=%s.%s&token=Y4119c3d7117497d680fae8277165a014", domain.Name, domain.Suffix)
					userurl := fmt.Sprintf("https://wanwang.aliyun.com/domain/searchresult/#/?keyword=%s&suffix=%s", domain.Name, domain.Suffix)
					domain.UserAddr = userurl
					domain.Addr = url
					ds = append(ds, domain)
				}
			}
		case 2:
			for i := 'a'; i <= 'z'; i++ {
				for j := 'a'; j <= 'z'; j++ {
					sum++
					domain := new(model.Domain)
					num := fmt.Sprintf("%c%c%s", i, j, "ex")
					domain.Name = num
					domain.Suffix = "top"
					domain.OrderID = strconv.Itoa(sum)
					url := fmt.Sprintf("https://checkapi.aliyun.com/check/checkdomain?domain=%s.%s&token=Y4119c3d7117497d680fae8277165a014", domain.Name, domain.Suffix)
					userurl := fmt.Sprintf("https://wanwang.aliyun.com/domain/searchresult/#/?keyword=%s&suffix=%s", domain.Name, domain.Suffix)
					domain.UserAddr = userurl
					domain.Addr = url
					ds = append(ds, domain)
				}
			}
		case 3:
			for i := 'a'; i <= 'z'; i++ {
				for j := 'a'; j <= 'z'; j++ {
					sum++
					domain := new(model.Domain)
					num := fmt.Sprintf("%c%c%s", i, j, "ex")
					domain.Name = num
					domain.Suffix = "pro"
					domain.OrderID = strconv.Itoa(sum)
					url := fmt.Sprintf("https://checkapi.aliyun.com/check/checkdomain?domain=%s.%s&token=Y4119c3d7117497d680fae8277165a014", domain.Name, domain.Suffix)
					userurl := fmt.Sprintf("https://wanwang.aliyun.com/domain/searchresult/#/?keyword=%s&suffix=%s", domain.Name, domain.Suffix)
					domain.UserAddr = userurl
					domain.Addr = url
					ds = append(ds, domain)
				}
			}
		}

	}
	return ds
}
func monitor() {
	//index的数量需要和爬虫的携程数量对应
	for index := 0; index < 16; index++ {
		<-model.ExitCh
	}
	close(model.ResultCh)
}
func main() {
	//1、获取2个字母的组合+后缀（.vip、.top、.io）。并放入*domain结构体,返回一个[]*domain结构体的切片
	ds := createString()
	//2、建立一个domain结构体的通道，把相关数据丢入通道
	domainChan := make(chan *model.Domain, 2100)
	for _, domain := range ds {
		domainChan <- domain
	}
	//3、用爬虫程序去爬数据并返回结果
	for i := 0; i < 16; i++ {
		go utils.Spider(domainChan)
	}
	close(domainChan)
	//4、开启监控Spider携程完成情况监控携程
	go monitor()
	//4、遍历结果写入excel文件
	dss := make([]*model.Domain, 0)
LOOP:
	for {
		select {
		case x, ok := <-model.ResultCh:
			if !ok {
				break LOOP
			}
			if !x.IsRegister {
				dss = append(dss, x)
			}
		default:
			time.Sleep(time.Millisecond * 200)
		}
	}
	utils.WriteXls(dss, "domain.xlsx")
}
