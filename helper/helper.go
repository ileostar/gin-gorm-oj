package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"net/smtp"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

// GetMd5 md5加密
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("gin-gorm-oj-key")

// GenerateToken 生成Token
func GenerateToken(identity, name string) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken 解析Token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// SendCode 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	// 发送者
	e.From = "Get <14709723891@163.com>"
	// 接收者
	e.To = []string{toUserEmail}
	// 标题
	e.Subject = "验证码发送测试"
	// 内容
	e.HTML = []byte("您的验证码是：<b>" + code + "</b>")
	// 这里的password不是邮箱密码，是授权码
	// err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "liuxinghao030@163.com", "IBWQRSIRGITJSOKT", "smtp.163.com"))

	// 返回 EOF 是，关闭SSL重试
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "14709723891@163.com", "IBWQRSIRGITJSOKT", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

}

// GetUUID 生成唯一标识
func GetUUID() string {
	return uuid.NewV4().String()
}
