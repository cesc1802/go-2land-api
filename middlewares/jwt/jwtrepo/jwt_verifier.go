package jwtrepo

import (
	"context"
	"go-rest-api/modules/user/userrepo"
)

type Verifier interface {
	String() string
	VerifyUser(ctx context.Context, userId string) bool
}

type jwtVerifier struct {
	uStore userrepo.FindUserByIdStorage
}

func NewVerifier(storage userrepo.FindUserByIdStorage) Verifier {
	return &jwtVerifier{
		uStore: storage,
	}
}

func (j *jwtVerifier) String() string {
	return "jwtMiddleware"
}

func (j *jwtVerifier) VerifyUser(ctx context.Context, userId string) bool {
	_, err := j.uStore.FindUserById(ctx, userId)
	return err == nil
}
