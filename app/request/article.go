package request

// Comment 评论
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// Article 文章
type Article struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	SpaceiD   int        `json:"space_id"`
	SpaceName string     `json:"space_name"`
	Like      int        `json:"like"`
	Comment   []*Comment `json:"comment"`
}

// ArticlePostRequest 发布文章 请求
type ArticlePostRequest struct {
	Title   string `json:"title" binding:"requird"`
	Content string `json:"content" binding:"requird"`
}

// ArticlePostResponse 发布文章 返回
type ArticlePostResponse struct {
}

// ArticleCommentRequest 发表评论 请求
type ArticleCommentRequest struct {
	Content string `json:"content"`
}

// ArticleCommentResponse 发表评论 返回
type ArticleCommentResponse struct {
}

// ArticleLikeRequest 点赞 请求
type ArticleLikeRequest struct {
}

// ArticleLikeResponse 点赞 返回
type ArticleLikeResponse struct {
}

// ArticleTopRequest 置顶 请求
type ArticleTopRequest struct {
}

// ArticleTopResponse 置顶 返回
type ArticleTopResponse struct {
}

// ArticleListRequest 列表 请求
type ArticleListRequest struct {
}

// ArticleListResponse 列表 返回
type ArticleListResponse struct {
	Items []*Article `json:"items"`
}

// ArticleDetailRequest 详情 请求
type ArticleDetailRequest struct {
}

// ArticleDetailResponse 详情 返回
type ArticleDetailResponse struct {
	Item *Article `json:"item"`
}
