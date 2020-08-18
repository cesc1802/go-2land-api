package postrepo

import (
	"context"
	"go-rest-api/modules/post/postmodel"
)

type GetPostStorage interface {
	GetAll(ctx context.Context) ([]postmodel.Post, error)
}

type getPostStorage struct {
	store GetPostStorage

}

func NewGetPostStorage(store GetPostStorage) *getPostStorage {
	return &getPostStorage{
		store: store,
	}
}

func (repo *getPostStorage) GetAll(ctx context.Context) ([]postmodel.Post, error) {
	var posts []postmodel.Post
	var err error

	posts, err = repo.store.GetAll(ctx)
	return posts, err
}
