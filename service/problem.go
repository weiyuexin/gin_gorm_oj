package service

import (
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
	models.GetProblemList()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "getProblemList",
		"data": "",
	})
}
