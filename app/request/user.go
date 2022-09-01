package request

// UserLoginRequest 登录 请求
type UserLoginRequest struct {
	Email  string `json:"email" binding:"requird"`
	Passwd string `json:"passwd" binding:"requird"`
}

// UserLoginResponse 登录 返回
type UserLoginResponse struct {
	Token string `json:"token"`
}

// UserRegisterRequest 注册 请求
type UserRegisterRequest struct {
	Email  string `json:"email" binding:"requird"`
	Passwd string `json:"passwd" binding:"requird"`
}

// UserRegisterResponse 注册 返回
type UserRegisterResponse struct {
}
