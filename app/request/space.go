package request

// SpaceCreateRequest 创建板块 请求
type SpaceCreateRequest struct {
	Name string `json:"name"`
}

// SpaceCreateResponse 创建板块 返回
type SpaceCreateResponse struct {
}

// SpaceJoinRequest 加入板块 请求
type SpaceJoinRequest struct {
	ID int64 `json:"id"`
}

// SpaceJoinResponse 加入板块 返回
type SpaceJoinResponse struct {
}

// SpaceAdminRequest 设置管理员 请求
type SpaceAdminRequest struct {
	ID  int64 `json:"id"`
	UID int64 `json:"uid"`
}

// SpaceAdminResponse 设置管理员 返回
type SpaceAdminResponse struct {
}
