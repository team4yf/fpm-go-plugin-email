package email

// Driver 邮件发送驱动接口定义
type Driver interface {
	// Send 发送邮件
	Send(to, subject, body string) error
	// Close 关闭链接
	Close()
}
