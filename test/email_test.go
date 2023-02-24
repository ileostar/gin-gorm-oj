package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	// 发送者
	e.From = "Get <14709723891@163.com>"
	// 接收者
	e.To = []string{"245956856@qq.com"}
	// 标题
	e.Subject = "验证码发送测试"
	// 内容
	e.HTML = []byte("您的验证码是：<b>123456</b>")
	// 这里的password不是邮箱密码，是授权码
	// err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "liuxinghao030@163.com", "IBWQRSIRGITJSOKT", "smtp.163.com"))

	// 返回 EOF 是，关闭SSL重试
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "14709723891@163.com", "IBWQRSIRGITJSOKT", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

	if err != nil {
		t.Fatal(err)
	}
}
