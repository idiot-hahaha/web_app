package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"web_app/models"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(*) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		//fmt.Printf("查询失败,err:%v\n", err)
		return
	}
	if count > 0 {
		//fmt.Printf("用户已存在\n")
		return errors.New("用户已存在")
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// 对用户密码进行加密
	user.Password = encryptPassword(user.Password)

	sqlStr := "insert into user (user_id, username, password) values(?,?,?) "

	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func Login(user *models.User) (err error) {
	sqlStr := "select user_id, username, password from user where username=?"
	oPassword := user.Password
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		zap.L().Error("登陆失败，用户不存在", zap.Error(err))
		return errors.New("用户不存在")
	}
	if err != nil {
		zap.L().Error("登陆信息查询失败", zap.Error(err))
		return
	}
	if user.Password != encryptPassword(oPassword) {
		err = errors.New("密码错误")
		zap.L().Error("登陆失败", zap.Error(err))
		return
	}
	return
}

const secret = "web_app"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
