package dto

// LoginReq登录请求参数
type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterReq 注册请求参数
type RegisterReq struct {
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	VerificationCode string `json:"verification_code" binding:"required"`
	Role             string `json:"role" binding:"required"`
}

// LoginResp 登录响应数据
type LoginResp struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"user_info"`
}

// UserInfo 用户基本信息
type UserInfo struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
}
