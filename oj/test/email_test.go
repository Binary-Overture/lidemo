package test

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%d,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, 123456)

	e := email.NewEmail()
	e.From = "Pornhub官方 <3398202776@qq.com>"
	e.To = []string{"2094458770@qq.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Pornhub 验证"
	e.Text = []byte("您的验证码是：")
	e.HTML = []byte(html)
	err := e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "3398202776@qq.com", "qlofsxjfsvvmcibi", "smtp.qq.com:465"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com:465"})
	if err != nil {
		return
	}
}
