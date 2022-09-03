package request

// UserLoginRequest 登录 请求
type UserLoginRequest struct {
	Email  string `json:"email" binding:"required"`
	Passwd string `json:"passwd" binding:"required"`
}

// UserLoginResponse 登录 返回
type UserLoginResponse struct {
	Token string `json:"token"`
}

// UserRegisterRequest 注册 请求
type UserRegisterRequest struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Passwd string `json:"passwd" binding:"required"`
}

// UserRegisterResponse 注册 返回
type UserRegisterResponse struct {
}
