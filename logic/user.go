package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// SignUp 用户注册
func SignUp(param *models.ParamSignUp) (err error) {

	//判断用户名是否被使用
	if err = mysql.CheckUserExist(param.Username); err != nil {
		return
	}

	// 通过雪花算法计算用户id
	userID := snowflake.GenID()

	// 创建用户实例
	user := &models.User{
		UserID:   userID,
		Username: param.Username,
		Password: param.Password,
	}

	// 将用户信息插入到表格中
	err = mysql.InsertUser(user)
	return
}

// Login 用户登录
func Login(p *models.ParamLogin) (err error) {
	// 创建用户实例
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//返回登陆结果
	return mysql.Login(user)
}
