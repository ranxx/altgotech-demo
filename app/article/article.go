package article

import (
	"context"
	"fmt"
	"time"

	"github.com/ranxx/altgotech-demo/app/request"
	"github.com/ranxx/altgotech-demo/pkg/db"
	"github.com/ranxx/altgotech-demo/service/article"
	"github.com/ranxx/altgotech-demo/service/article/model"
)

// Article 文章
type Article struct {
}

// NewArticle new
func NewArticle(ctx context.Context) *Article {
	return &Article{}
}

// GetService new svc
func (a *Article) GetService() *article.Service {
	return article.NewService()
}

// Post 发布文章
func (a *Article) Post(ctx context.Context, uid int, req *request.ArticlePostRequest) (*request.ArticlePostResponse, error) {
	svc := a.GetService()

	item := &model.Article{
		Title:     req.Title,
		Content:   req.Content,
		Tags:      req.Tags,
		SpaceID:   req.Space,
		CreatoriD: uid,
	}

	err := svc.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return &request.ArticlePostResponse{
		Item: ConvertArticleFromModel(item),
	}, nil
}

// ConvertCommentFromModel convert
func ConvertCommentFromModel(item *model.Comment) *request.Comment {
	return &request.Comment{
		ID:      item.ID,
		PID:     item.CID,
		AID:     item.AID,
		Content: item.Content,
	}
}

// Comment 评论
func (a *Article) Comment(ctx context.Context, uid int, req *request.ArticleCommentRequest) (*request.ArticleCommentResponse, error) {
	svc := a.GetService()

	item := &model.Comment{
		Content: req.Content,
		AID:     req.AID,
		UID:     uid,
		CID:     req.PID,
	}

	if err := svc.Comment(ctx, item); err != nil {
		return nil, err
	}

	return &request.ArticleCommentResponse{
		Item: ConvertCommentFromModel(item),
	}, nil
}

// Like 点赞
func (a *Article) Like(ctx context.Context, uid int, req *request.ArticleLikeRequest) (*request.ArticleLikeResponse, error) {
	svc := a.GetService()

	key := fmt.Sprintf("Blog:Article:Like:%d:%d", req.AID, uid)

	// 是否已经点赞
	exist, _ := db.Redis().Exists(ctx, key).Result()
	if exist == 1 {
		return &request.ArticleLikeResponse{}, nil
	}

	// 先查询是否有
	ok, err := svc.Liked(ctx, uid, req.AID)
	if err != nil {
		return nil, err
	}
	if ok {
		// 15天
		db.Redis().Set(ctx, key, 1, time.Hour*24*15)
		return &request.ArticleLikeResponse{}, nil
	}

	if err := svc.Like(ctx, uid, req.AID); err != nil {
		return nil, err
	}

	return &request.ArticleLikeResponse{}, nil
}

// UnLike 取消点赞
func (a *Article) UnLike(ctx context.Context, uid int, req *request.ArticleUnLikeRequest) (*request.ArticleUnLikeResponse, error) {
	svc := a.GetService()

	key := fmt.Sprintf("Blog:Article:Like:%d:%d", req.AID, uid)

	svc.UnLike(ctx, uid, req.AID)
	db.Redis().Del(ctx, key)

	return &request.ArticleUnLikeResponse{}, nil
}

// Top 置顶
func (a *Article) Top(ctx context.Context, uid int, req *request.ArticleTopRequest) (*request.ArticleTopResponse, error) {
	svc := a.GetService()

	err := svc.Top(ctx, req.AID, req.Top)

	return &request.ArticleTopResponse{}, err
}

// ConvertArticleFromModel convert
func ConvertArticleFromModel(item *model.Article) *request.Article {
	return &request.Article{
		ID:      item.ID,
		Title:   item.Title,
		Content: item.Content,
		Tags:    item.Tags,
		SpaceiD: item.SpaceID,
		Comment: []*request.Comment{},
		Top:     item.Top,
	}
}

// List 文章列表
func (a *Article) List(ctx context.Context, uid int, req *request.ArticleListRequest) (*request.ArticleListResponse, error) {
	svc := a.GetService()

	items, total, err := svc.List(ctx, req.Space, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	// 点赞数和评论数
	aids := make([]int, 0, len(items))
	for _, v := range items {
		aids = append(aids, v.ID)
	}

	likeMap, err := svc.BatchLikeTotal(ctx, aids...)
	if err != nil {
		likeMap = map[int]int64{}
	}
	commentMap, err := svc.BatchCommentTotal(ctx, aids...)
	if err != nil {
		commentMap = map[int]int64{}
	}

	resp := &request.ArticleListResponse{
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  total,
		Items:  make([]*request.Article, 0, len(items)),
	}

	for _, v := range items {
		item := ConvertArticleFromModel(v)
		item.LikeTotal = likeMap[v.ID]
		item.CommentTotal = commentMap[v.ID]
		resp.Items = append(resp.Items, item)
	}

	return resp, err
}

// Detail 文章详情
func (a *Article) Detail(ctx context.Context, uid int, req *request.ArticleDetailRequest) (*request.ArticleDetailResponse, error) {
	svc := a.GetService()

	item, err := svc.Get(ctx, req.AID)
	if err != nil {
		return nil, err
	}

	like, _ := svc.LikeTotal(ctx, req.AID)
	comments, _, _ := svc.ListComment(ctx, int64(req.AID), -1, -1) // -1 全部返回
	ok, _ := svc.Liked(ctx, uid, req.AID)

	article := ConvertArticleFromModel(item)
	article.Comment = make([]*request.Comment, 0, len(comments))
	article.LikeTotal = like
	article.Liked = ok
	article.CommentTotal = int64(len(comments))

	for _, v := range comments {
		article.Comment = append(article.Comment, ConvertCommentFromModel(v))
	}

	return &request.ArticleDetailResponse{
		Item: article,
	}, nil
}
