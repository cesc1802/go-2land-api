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


	//if err := db.Table(postmodel.Post{}.TableName()).Find(&posts).Error; err != nil {
	//	db.Rollback()
	//	return nil, nil
	//}
	//
	//if err := db.Commit().Error; err != nil {
	//	db.Rollback()
	//	return nil, errors.New("cannot create new a post")
	//}
	return posts, nil
}
