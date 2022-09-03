package dao

import (
	"context"

	"github.com/ranxx/altgotech-demo/pkg/db"
	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/service/article/model"
	"gorm.io/gorm"
)

// Article 文章
type Article struct {
	db *gorm.DB
}

// NewArticle new article
func NewArticle(db *gorm.DB) *Article {
	return &Article{
		db: db,
	}
}

// Create 创建
func (a *Article) Create(ctx context.Context, item *model.Article) error {
	return a.db.WithContext(ctx).Create(item).Error
}

// Comment 评论
func (a *Article) Comment(ctx context.Context, item *model.Comment) error {
	return a.db.WithContext(ctx).Create(item).Error
}

// DelComment 删除评论
func (a *Article) DelComment(ctx context.Context, id int) error {
	return a.db.Model(&model.Comment{}).WithContext(ctx).Where("id = ?", id).Delete(nil).Error
}

// Like 点赞
func (a *Article) Like(ctx context.Context, item *model.Like) error {
	return a.db.WithContext(ctx).Create(item).Error
}

// UnLike 取消点赞
func (a *Article) UnLike(ctx context.Context, uid, aid int) error {
	return a.db.Model(&model.Like{}).WithContext(ctx).Where("uid = ? and aid = ?", uid, aid).Delete(nil).Error
}

// Top 置顶
func (a *Article) Top(ctx context.Context, id int, top bool) error {
	return a.db.Model(&model.Article{}).Where("id = ?", id).Update("top", top).Error
}

// Get 获取
func (a *Article) Get(ctx context.Context, id int) (*model.Article, error) {
	item := model.Article{}

	ret := a.db.Model(&model.Article{}).WithContext(ctx).Where("id = ?", id).First(&item)
	if ret.Error != nil {
		return nil, ret.Error
	}

	if ret.RowsAffected == 0 {
		return nil, errors.NewErrCode(errors.ErrArticleNotFound, "文章不存在")
	}

	return &item, nil
}

// List 列表
func (a *Article) List(ctx context.Context, space int, offset, limit int64) ([]*model.Article, int64, error) {
	items := []*model.Article{}
	total := int64(0)

	dbr := a.db.Model(&model.Article{}).WithContext(ctx).Where("space_id = ?", space).Order("top desc,created_at desc")

	if err := dbr.Count(&total).Error; err != nil || total <= 0 {
		return nil, 0, err
	}

	if err := db.Pager(dbr, offset, limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// ListComment 列表评论
func (a *Article) ListComment(ctx context.Context, aid, offset, limit int64) ([]*model.Comment, int64, error) {
	items := []*model.Comment{}
	total := int64(0)

	dbr := a.db.Model(&model.Comment{}).WithContext(ctx).Where("aid = ?", aid)

	if err := dbr.Count(&total).Error; err != nil || total <= 0 {
		return nil, 0, err
	}

	if err := db.Pager(dbr, offset, limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// Liked 是否点赞
func (a *Article) Liked(ctx context.Context, uid, aid int) (bool, error) {
	total := (int64)(0)

	err := a.db.Model(&model.Like{}).WithContext(ctx).Where("uid = ? and aid = ?", uid, aid).Count(&total).Error

	return total > 0, err
}

// LikeTotal 点赞数
func (a *Article) LikeTotal(ctx context.Context, aid int) (int64, error) {
	total := (int64)(0)

	err := a.db.Model(&model.Like{}).WithContext(ctx).Where("aid = ?", aid).Count(&total).Error

	return total, err
}

// BatchLikeTotal 批量获取点赞数
func (a *Article) BatchLikeTotal(ctx context.Context, aid ...int) ([]*model.Total, error) {
	totals := []*model.Total{}
	if len(aid) <= 0 {
		return totals, nil
	}

	err := a.db.Model(&model.Like{}).Select("aid as id,count(1) as total").Where("aid in (?)", aid).Group("aid").Find(&totals).Error

	return totals, err
}

// CommentTotal 评论数
func (a *Article) CommentTotal(ctx context.Context, aid int) (int64, error) {
	total := (int64)(0)

	err := a.db.Model(&model.Comment{}).WithContext(ctx).Where("aid = ?", aid).Count(&total).Error

	return total, err
}

// BatchCommentTotal 批量获取评论数
func (a *Article) BatchCommentTotal(ctx context.Context, aid ...int) ([]*model.Total, error) {
	totals := []*model.Total{}
	if len(aid) <= 0 {
		return totals, nil
	}

	err := a.db.Model(&model.Comment{}).Select("aid as id,count(1) as total").Where("aid in (?)", aid).Group("aid").Find(&totals).Error

	return totals, err
}
