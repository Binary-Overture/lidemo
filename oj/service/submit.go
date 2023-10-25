package service

import (
	"cncyx.xyz/define"
	"cncyx.xyz/models"
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
// @Router /problem-submit [get]
func GetSubmitList(c *gin.Context) {
	var count int64
	data := make([]*models.SubmitBasic, 0)
	offset, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetSubmitList page get error :", err)
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetSubmitList page get error :", err)
	}
	page := (offset - 1) * size
	problemIdentity := c.Query("problem-identity")
	userIdentity := c.Query("user-identity")

	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		log.Println("GetSubmitList problem status get error :", err)
	}
	tx := models.GetSubmitListByAll(problemIdentity, userIdentity, status)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetSubmitList error :" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"count": count,
			"list":  data,
		},
	})
}
