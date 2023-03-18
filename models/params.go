package models

// ParamSignUp 用户注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Age        uint8  `json:"age" binding:"gte=1,lte=130" `
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 用户登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
