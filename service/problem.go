package service

import (
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags Problem
// @Summary 获取题目列表
// @Schemes
// @Description 获取题目列表
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query int false "keyword"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	//读取参数
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("Get Problem List strconv Error : ", err)
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")

	list := make([]*models.Problem, 0)
	//获取keyword查询到的数据
	tx := models.GetProblemList(keyword)
	//分页
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error : ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
