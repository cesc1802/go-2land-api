package poststorage

import (
	"context"
	"errors"
	"go-rest-api/modules/post/postmodel"
)


func (p *postSQLStorage) Create(ctx context.Context, post *postmodel.Post) error {
	db := p.db.New().Begin()

	if err := db.Table(postmodel.Post{}.TableName()).Create(post); err != nil {
		db.Rollback()
		return errors.New("cannot create new a post")
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return errors.New("cannot create new a post")
	}
	return nil
}
