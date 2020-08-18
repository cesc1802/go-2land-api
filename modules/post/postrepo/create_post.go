package postrepo

import (
	"context"
	"go-rest-api/modules/post/postmodel"
)

type CreatePostStorage interface {
	Create(ctx context.Context, postData *postmodel.Post) error
}

type createPost struct {
	store CreatePostStorage
}

func NewCreatePost(store CreatePostStorage) *createPost {
	return &createPost{store: store}
}

func (cr *createPost) CreatePost(ctx context.Context, data *postmodel.Post) error {
	return cr.store.Create(ctx, data)
}
