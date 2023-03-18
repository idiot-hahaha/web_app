package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/models"
)

func SignUpHandler(c *gin.Context) {

	// 1.参数校验
	param := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("注册信息错误", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册信息错误" + err.Error(),
		})
		return
	}

	// 2.业务处理
	if err := logic.SignUp(param); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败:" + err.Error(),
		})
		return
	}

	// 3.注册成功后返回响应
	c.JSON(http.StatusOK, "注册成功")
	return
}

func LoginHandler(c *gin.Context) {
	// 1.校验参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("登录信息校验错误", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录校验错误" + err.Error(),
		})
		return
	}

	//	2.登录功能业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("登录信息错误", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登陆失败:" + err.Error(),
		})
		return
	}

	//	3.登陆成功
	c.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}
