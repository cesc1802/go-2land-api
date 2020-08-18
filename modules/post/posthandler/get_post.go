package posthandler

import (
	"context"
	"go-rest-api/modules/post/postmodel"
)

type GetPostHdl interface {
	GetAll(ctx context.Context) ([] postmodel.Post, error)
}

type getPostHdl struct {
	repo GetPostHdl
}

func NewGetPostHdl(repo GetPostHdl) *getPostHdl {
	return &getPostHdl{
		repo: repo,
	}
}

func (gph *getPostHdl) GetAll(ctx context.Context) ([] postmodel.Post, error) {
	return gph.repo.GetAll(ctx)
}
