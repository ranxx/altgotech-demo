package article

import (
	"context"

	"github.com/ranxx/altgotech-demo/pkg/db"
	"github.com/ranxx/altgotech-demo/service/article/dao"
	"github.com/ranxx/altgotech-demo/service/article/model"
)

// Service article
type Service struct {
	d *dao.Article
}

// NewService new
func NewService() *Service {
	return &Service{
		d: dao.NewArticle(db.Mysql()),
	}
}

// Create 创建
func (s *Service) Create(ctx context.Context, item *model.Article) error {
	return s.d.Create(ctx, item)
}

// Comment 评论
func (s *Service) Comment(ctx context.Context, item *model.Comment) error {
	return s.d.Comment(ctx, item)
}

// DelComment 删除评论
func (s *Service) DelComment(ctx context.Context, id int) error {
	return s.d.DelComment(ctx, id)
}

// Like 点赞
func (s *Service) Like(ctx context.Context, uid, aid int) error {
	return s.d.Like(ctx, &model.Like{
		AID: aid,
		UID: uid,
	})
}

// UnLike 取消点赞
func (s *Service) UnLike(ctx context.Context, uid, aid int) error {
	return s.d.UnLike(ctx, uid, aid)
}

// Top 置顶
func (s *Service) Top(ctx context.Context, id int, top bool) error {
	return s.d.Top(ctx, id, top)
}

// Get 获取
func (s *Service) Get(ctx context.Context, id int) (*model.Article, error) {
	return s.d.Get(ctx, id)
}

// List 列表
func (s *Service) List(ctx context.Context, space int, offset, limit int64) ([]*model.Article, int64, error) {
	return s.d.List(ctx, space, offset, limit)
}

// ListComment 列表评论
func (s *Service) ListComment(ctx context.Context, aid, offset, limit int64) ([]*model.Comment, int64, error) {
	return s.d.ListComment(ctx, aid, offset, limit)
}

// Liked 是否点赞
func (s *Service) Liked(ctx context.Context, uid, aid int) (bool, error) {
	return s.d.Liked(ctx, uid, aid)
}

// LikeTotal 点赞数
func (s *Service) LikeTotal(ctx context.Context, aid int) (int64, error) {
	return s.d.LikeTotal(ctx, aid)
}

// BatchLikeTotal 批量获取点赞数
func (s *Service) BatchLikeTotal(ctx context.Context, aid ...int) (map[int]int64, error) {
	items, err := s.d.BatchLikeTotal(ctx, aid...)
	if err != nil {
		return map[int]int64{}, err
	}
	ret := map[int]int64{}
	for _, v := range items {
		ret[v.ID] = v.Total
	}
	return ret, nil
}

// CommentTotal 评论数
func (s *Service) CommentTotal(ctx context.Context, aid int) (int64, error) {
	return s.d.CommentTotal(ctx, aid)
}

// BatchCommentTotal 批量获取评论数
func (s *Service) BatchCommentTotal(ctx context.Context, aid ...int) (map[int]int64, error) {
	items, err := s.d.BatchCommentTotal(ctx, aid...)
	if err != nil {
		return map[int]int64{}, err
	}
	ret := map[int]int64{}
	for _, v := range items {
		ret[v.ID] = v.Total
	}
	return ret, nil
}
