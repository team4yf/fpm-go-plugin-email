package plugin

import (
	"github.com/team4yf/fpm-go-plugin-email/email"
	"github.com/team4yf/yf-fpm-server-go/fpm"
)

type sendReq struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func init() {
	fpm.RegisterByPlugin(&fpm.Plugin{
		Name: "fpm-plugin-email",
		V:    "0.0.1",
		Handler: func(fpmApp *fpm.Fpm) {
			smtpConfig := email.SMTPConfig{}
			if fpmApp.HasConfig("smtp") {
				if err := fpmApp.FetchConfig("smtp", &smtpConfig); err != nil {
					panic(err)
				}
			}

			fpmApp.Logger.Debugf("Startup : %s, config: %v", "email", smtpConfig)

			client := email.NewSMTPClient(smtpConfig)
			client.Init()

			fpmApp.AddBizModule("email", &fpm.BizModule{
				"send": func(param *fpm.BizParam) (data interface{}, err error) {
					var req sendReq
					if err = param.Convert(&req); err != nil {
						return
					}
					err = client.Send(req.To, req.Subject, req.Body)
					data = 1
					return
				},
			})

		},
	})
}
