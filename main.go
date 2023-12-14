package main

import (
	"fmt"
	"github.com/1152545264/goSpyder/collect"
	"github.com/1152545264/goSpyder/log"
	"github.com/1152545264/goSpyder/parse/doubangroup"
	"github.com/1152545264/goSpyder/proxy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	// log
	plugin := log.NewStdoutPlugin(zapcore.InfoLevel)
	logger := log.NewLogger(plugin)
	logger.Info("log init end")

	// proxy
	proxyURLs := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8888"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		logger.Error("RoundRobinProxySwitcher failed")
	}

	// douban cookie
	cookie := "bid=i5EpPdK4zoM; ct=y; viewed=\"25743846_1168618\"; _pk_id.100001.8cb4=aadcf3aa5ed2b084.1702529147.; _pk_ses.100001.8cb4=1; ap_v=0,6.0; dbcl2=\"172026410:ax1Vpu5LWhc\"; ck=AeAa; push_noty_num=0; push_doumail_num=0"

	var worklist []*collect.Request
	for i := 0; i <= 0; i += 25 {
		str := fmt.Sprintf("https://www.douban.com/group/szsh/discussion?start=%d", i)
		worklist = append(worklist, &collect.Request{
			Url:       str,
			Cookie:    cookie,
			ParseFunc: doubangroup.ParseURL,
		})
	}

	var f collect.Fetcher = &collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			body, err := f.Get(item)
			time.Sleep(1 * time.Second)
			if err != nil {
				logger.Error("read content failed",
					zap.Error(err),
				)
				continue
			}
			res := item.ParseFunc(body, item)
			for _, item := range res.Items {
				logger.Info("result",
					zap.String("get url:", item.(string)))
			}
			worklist = append(worklist, res.Requests...)
		}
	}

}
