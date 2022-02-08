package controller

import (
	"blue_me/logic"
	"blue_me/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应,可以记录个日志zap
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求的参数有误",
		})
		return
	}
	fmt.Println(p)
	// 2.业务处理
	logic.SignUp()
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
