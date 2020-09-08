package main

import (
	"github.com/team4yf/yf-fpm-server-go/fpm"

	_ "github.com/team4yf/fpm-go-plugin-email/plugin"
)

func main() {
	fpmApp := fpm.New()

	fpmApp.Init()

	fpmApp.Execute("email.send", &fpm.BizParam{
		"to":      "1794947912@qq.com",
		"subject": "test",
		"body":    "ok",
	})

	fpmApp.Run()
}
