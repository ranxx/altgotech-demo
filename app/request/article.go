package request

// Comment 评论
type Comment struct {
	ID      int    `json:"id"`
	PID     int    `json:"pid"`
	AID     int    `json:"aid"`
	Content string `json:"content"`
}

// Article 文章
type Article struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	Tags         string     `json:"tags"`
	SpaceiD      int        `json:"space_id"`
	SpaceName    string     `json:"space_name"`
	Comment      []*Comment `json:"comment"`
	LikeTotal    int64      `json:"like_total"`
	CommentTotal int64      `json:"comment_total"`
	Liked        bool       `json:"liked"`
	Top          bool       `json:"top"`
}

// ArticlePostRequest 发布文章 请求
type ArticlePostRequest struct {
	Title   string `json:"title" binding:"gte=1"`
	Content string `json:"content" binding:"gte=1"`
	Tags    string `json:"tags"`
	Space   int    `json:"space"`
}

// ArticlePostResponse 发布文章 返回
type ArticlePostResponse struct {
	Item *Article `json:"item"`
}

// ArticleCommentRequest 发表评论 请求
type ArticleCommentRequest struct {
	AID     int    `json:"aid"`
	PID     int    `json:"pid"`
	Content string `json:"content" binding:"gte=2"`
}

// ArticleCommentResponse 发表评论 返回
type ArticleCommentResponse struct {
	Item *Comment `json:"item"`
}

// ArticleLikeRequest 点赞 请求
type ArticleLikeRequest struct {
	AID int `json:"aid"`
}

// ArticleLikeResponse 点赞 返回
type ArticleLikeResponse struct {
}

// ArticleUnLikeRequest 取消点赞 请求
type ArticleUnLikeRequest struct {
	AID int `json:"aid"`
}

// ArticleUnLikeResponse 取消点赞 返回
type ArticleUnLikeResponse struct {
}

// ArticleTopRequest 置顶 请求
type ArticleTopRequest struct {
	AID   int  `json:"aid"`
	Space int  `json:"space"`
	Top   bool `json:"top"`
}

// ArticleTopResponse 置顶 返回
type ArticleTopResponse struct {
}

// ArticleListRequest 列表 请求
type ArticleListRequest struct {
	Space  int   `json:"space"`
	Offset int64 `json:"offset" form:"offset"`
	Limit  int64 `json:"limit" form:"limit"`
}

// ArticleListResponse 列表 返回
type ArticleListResponse struct {
	Offset int64      `json:"offset"`
	Limit  int64      `json:"limit"`
	Total  int64      `json:"total"`
	Items  []*Article `json:"items"`
}

// ArticleDetailRequest 详情 请求
type ArticleDetailRequest struct {
	AID   int `json:"aid"`
	Space int `json:"space"`
}

// ArticleDetailResponse 详情 返回
type ArticleDetailResponse struct {
	Item *Article `json:"item"`
}
