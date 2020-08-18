package posthandler

import (
	"context"
	"go-rest-api/modules/post/postmodel"
)

type CreatePostHdl interface {
	CreatePost(ctx context.Context, postData *postmodel.Post) error
}

type createPostHdl struct {
	repo CreatePostHdl
}

func NewCreatePostHdl(repo CreatePostHdl) *createPostHdl {
	return &createPostHdl{
		repo: repo,
	}
}

func (cph *createPostHdl) CreatePost(ctx context.Context, postData *postmodel.Post) error {
	return cph.repo.CreatePost(ctx, postData)
}
