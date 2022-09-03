package request

// Space ...
type Space struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatorID int    `json:"creator_id"`
	AdminID   int    `json:"admin_id"`
}

// SpaceCreateRequest 创建板块 请求
type SpaceCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

// SpaceCreateResponse 创建板块 返回
type SpaceCreateResponse struct {
	Item *Space `json:"item"`
}

// SpaceListRequest 板块列表 请求
type SpaceListRequest struct {
	All bool `json:"all" form:"all"` // 传 true 则展示全部
}

// SpaceListResponse 板块列表 返回
type SpaceListResponse struct {
	Items []*Space `json:"items"`
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
	UID int64 `json:"uid" binding:"gte=1"`
}

// SpaceAdminResponse 设置管理员 返回
type SpaceAdminResponse struct {
}
