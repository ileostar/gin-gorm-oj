package service

import (
	"getcahrzp.cn/define"
	"getcahrzp.cn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param problem_identity query string false "problem_identity"
// @Param user_identity query string false "user_identity"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
	//给size和page赋默认值
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	// page == 1 ==> offset 0
	var count int64
	list := make([]*models.SubmitBasic, 0)

	//获取参数
	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	//调用查询方法
	tx := models.GetSubmitList(problemIdentity, userIdentity, status)

	//分页
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Problem List Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"data": map[string]interface{}{
			"data":  list,
			"count": count,
		},
	})

}
