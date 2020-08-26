package auth

import (
	"context"
	"errors"
	"time"
)

const (
	BeaereScheme = "Bearer "
	ScopePublic  = ""
	ScopeAccount = "*"
)

var (
	ErrInvalidToken = errors.New("invalid token provided")
	ErrForbidden    = errors.New("resource forbidden")
)

type Auth interface {
	// Init the auth
	Init(opts ...Option)
	// Options set for auth
	Options() Options
	// Generate a new account
	Generate(id string, opts ...GenerateOption) (*Account, error)
	// Verify an account has access to a resource using the rules
	Verify(acc *Account, res *Resource, opts ...VerifyOption) error
	// Inspect a token
	Inspect(token string) (*Account, error)
	// Token generated using refresh token or credentials
	Token(opts ...TokenOption) (*Token, error)
	// Grant access to a resource
	Grant(rule *Rule) error
	// Revoke access to a resource
	Revoke(rule *Rule) error
	// Rules returns all the rules used to verify requests
	Rules(...RulesOption) ([]*Rule, error)
	// String returns the name of the implementation
	String() string
}

type Account struct {
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	Issuer   string            `json:"issuer"`
	Metadata map[string]string `json:"metadata"`
	Scopes   []string          `json:"scopes"`
	Secret   string            `json:"secret"`
}

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Created      time.Time `json:"created"`
	Expiry       time.Time `json:"expiry"`
}

type Resource struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}

type Access int

const (
	AccessGranted Access = iota
	AccessDenied
)

type Rule struct {
	ID       string
	Scope    string
	Resource *Resource
	Access   Access
	Priority int32
}

type accountKey struct {
}

func AccountFromContext(ctx context.Context) (*Account, bool) {
	acc, ok := ctx.Value(accountKey{}).(*Account)
	return acc, ok
}

func ContextWithAccount(ctx context.Context, account *Account) context.Context {
	return context.WithValue(ctx, accountKey{}, account)
}
