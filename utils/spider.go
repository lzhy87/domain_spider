package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lzhy87/domain_spider/model"
)

//Spider 爬取网页内容2
func Spider(domainChan <-chan *model.Domain) {
	for {
		select {
		//如果spiderChan通道关了，就获取不到数据了，就退出了
		case x, ok := <-domainChan:
			if !ok {
				//发送一个bool值至model.ExitCh管道（这里开几个携程就会发送几次bool值)
				model.ExitCh <- true
				fmt.Println("我获取不了数据了，saygoodbye!!")
				return
			}
			// 执行爬虫程序
			// time.Sleep(time.Millisecond * 600)
			// fmt.Println(x.Addr)
			req, err := http.Get(x.Addr)
			if err != nil {
				fmt.Println("http.Get err=", err)
				return
			}
			req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
			defer req.Body.Close()
			// f, _ := os.OpenFile("1.html", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

			buf, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Println("http read error", err)
				return
			}
			domain := new(model.AliyunDomain)
			//通过接口获取域名信息并反序列化至model.AliyunDomain结构体
			json.Unmarshal(buf, domain)
			//循环遍历结构体判断结构体内的avail字段是否等于1，如果等于0就把Domain结构的IsRegister字段设置为true
			//v.Avail=0代表不可用，=1代表可用
			for _, v := range domain.Module {
				if v.Avail == 0 {
					x.IsRegister = true
				}

			}
			model.ResultCh <- x
			fmt.Println(domain)
			fmt.Println(x.IsRegister)
			// reg := regexp.MustCompile(".+\\d\\s")

			// str1 := strings.TrimSpace(reg.FindString(str))
			// fmt.Printf("%#v", str1)

			// x.Content = str
			//把ETE结构体发送至model.Ch管道

			// fmt.Println("Spider:", x.OrderID, "  ", x.Name)

		// case <-ctx.Done():
		// 	return
		default:
			fmt.Println("Spider正在休眠")
			time.Sleep(time.Millisecond * 30)
		}

	}
}

//FilterData 数据清理
// func FilterData(ctx context.Context, symbol string) {
// 	for {
// 		select {
// 		//从model.Ch管道中获取ete结构体的数据
// 		case order, ok := <-model.Ch:
// 			//处理获取到的ete结构体数据
// 			if !ok {
// 				return
// 			}
// 			reg := regexp.MustCompile(".+\\d\\s" + symbol)

// 			str := strings.TrimSpace(reg.FindString(order.Content))
// 			if str == "" {
// 				fmt.Println("FilterData:", order.OrderID, "+", order.Name, "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
// 			}
// 			order.Balance = str
// 			fmt.Println(str)
// 			//把处理好的ete结构体数据发送至model.ResultCh管道
// 			model.ResultCh <- order
// 		case <-ctx.Done():
// 			return
// 		default:

// 			time.Sleep(time.Millisecond * 30)
// 		}

// 	}

// }
