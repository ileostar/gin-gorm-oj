package service

import (
	"getcahrzp.cn/helper"
	"getcahrzp.cn/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// GetUserDetail
// @Tags 公共方法
// @Summary 用户详情
// @Param identity query string false "user identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识不能为空",
		})
		return
	}

	data := new(models.UserBasic)
	// 通过Omit()省略字段
	err := models.Db.Omit("password").Where("identity= ?", identity).Find(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "当前用户标识不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserDetail Error:" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}

	// md5 加密
	password = helper.GetMd5(password)
	print(username, password)

	data := new(*models.UserBasic)
	err := models.Db.Where("name=? AND password = ?", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msd":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msd":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}

	// 获取token
	token, err := helper.GenerateToken((*data).Identity, (*data).Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msd":  "GenerateToken Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"token": token,
		},
	})
}

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	code := "123456"
	models.Rdb.Set(c, email, code, time.Second*300)
	err := helper.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send Code Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param mail formData string true "mail"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	mail := c.PostForm("mail")
	userCode := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if mail == "" || userCode == "" || name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}

	// 判断验证码是否正确
	sysCode, err := models.Rdb.Get(c, mail).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Code Error:" + err.Error(),
		})
		return
	}
	if sysCode != userCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	// 数据插入
	userIdentity := helper.GetUUID()
	data := &models.UserBasic{
		Identity: userIdentity,
		Name:     name,
		Password: helper.GetMd5(password),
		Phone:    phone,
		Mail:     mail,
	}
	err = models.Db.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create User Error:" + err.Error(),
		})
		return
	}

	// 生成Token
	token, err := helper.GenerateToken(userIdentity, name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Token Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"token": token,
		},
	})
}
