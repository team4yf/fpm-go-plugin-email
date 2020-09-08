## fpm-go-plugin-email

### Install

`$ go get -u github.com/team4yf/fpm-go-plugin-email`

```golang

import _ "github.com/team4yf/fpm-go-plugin-email/plugin"

```

### Config
```yaml
smtp:
    host: smtp.qq.com
    port: 587            # 端口
    username: 2917071792@qq.com
    password: xxx # 这里不要写密码，密码在环境变量中保存 FPM_EMAIL_PASSWORD
    name: no_reply
    address: 2917071792@qq.com
    reply: 2917071792@qq.com
    keepalive: 30         # 连接保持时长
```

### Usage

```golang

fpmApp.Execute("email.send", &fpm.BizParam{
    "to":      "1794947912@qq.com",
    "subject": "test",
    "body":    "ok",
})

```