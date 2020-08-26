package poststorage

import (
	"context"
	"errors"
	"fmt"
	"go-rest-api/modules/post/postmodel"
)

func (p *postSQLStorage) GetAll(ctx context.Context) ([]postmodel.Post, error) {

	var posts []postmodel.Post
	db := p.db.New().Begin()

	if err := db.Select("*").Preload("Comments").Find(&posts).Error; err != nil {
		fmt.Println(err.Error())
		db.Rollback()
		return nil, errors.New("Cannot get all post")
	}

	return posts, nil
}
