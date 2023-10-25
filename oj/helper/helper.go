package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var key = []byte("gorm-gin-oj-key")

// GenerateToken
// 生成Token
func GenerateToken(identity, name string) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	stringToken, err := token.SignedString(key)
	fmt.Println(identity + "-------" + name)
	fmt.Println(stringToken)
	if err != nil {
		return "", err
	}
	return stringToken, nil
}

// AnalyseToken
// 解析Token
func AnalyseToken(token string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(token, userClaim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		fmt.Println(userClaim)
		return nil, err
	}
	return userClaim, nil
}

// GetMd5
// 密码进行md5转换
func GetMd5(c string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(c)))
}

// SendEmail
// 发送邮件，6位验证码
func SendEmail(toUserEmail string, code int) error {
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
    </div>`, code)

	e := email.NewEmail()
	e.From = "Pornhub官方 <3398202776@qq.com>"
	e.To = []string{toUserEmail}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Pornhub 验证"
	e.Text = []byte("您的验证码是：")
	e.HTML = []byte(html)
	return e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "3398202776@qq.com", "qlofsxjfsvvmcibi", "smtp.qq.com:465"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com:465"})
}

func GetRandomNumber() int {
	// 设置随机种子，以确保每次运行都有不同的随机数
	rand.NewSource(time.Now().UnixNano())
	// 生成六位随机数
	return rand.Intn(900000) + 100000
}

// GetUUID
// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}
