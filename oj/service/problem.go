package service

import (
	"cncyx.xyz/define"
	"cncyx.xyz/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problemTest [get]
func GetProblemList(c *gin.Context) {
	var count int64
	list := make([]*models.ProblemBasic, 0)
	offset, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList page get error :", err)
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetProblemList page get error :", err)
	}
	page := (offset - 1) * size
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	tx := models.GetProblemList(keyword, categoryIdentity)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetProblemList query error :", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":  list,
			"count": count,
		},
	})
}

// GetProblemDetail
// @Tags 公共方法
// @Summary 问题列表
// @Param identity query string false "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemDetail(c *gin.Context) {
	//根据identity唯一标识查找问题内容
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题标识为空",
		})
		return
	}
	data := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).Preload("ProblemCategory").Preload("ProblemCategory.CategoryBasic").Find(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "未查询到问题记录",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Have Problem Error :" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": data,
	})

}
