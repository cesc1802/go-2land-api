package token

import (
	"time"
)

type Options struct {
	// PathToPublicKey base64 encoded, used by JWT
	PathToPublicKey string
	// PathToPrivateKey base64 encoded, used by JWT
	PathToPrivateKey string
}

type Option func(o *Options)

// WithStore sets the token providers store
//func WithStore(s store.Store) Option {
//	return func(o *Options) {
//		o.Store = s
//	}
//}

// WithPathToPublicKey sets the JWT public key
func WithPathToPublicKey(key string) Option {
	return func(o *Options) {
		o.PathToPublicKey = key
	}
}

// WithPathToPrivateKey sets the JWT private key
func WithPathToPrivateKey(key string) Option {
	return func(o *Options) {
		o.PathToPrivateKey = key
	}
}

func NewOptions(opts ...Option) Options {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	//set default store
	//if options.Store == nil {
	//	options.Store = store.DefaultStore
	//}
	return options
}

type GenerateOptions struct {
	// Expiry for the token
	Expiry time.Duration
}

type GenerateOption func(o *GenerateOptions)

// WithExpiry for the generated account's token expires
func WithExpiry(d time.Duration) GenerateOption {
	return func(o *GenerateOptions) {
		o.Expiry = d
	}
}

// NewGenerateOptions from a slice of options
func NewGenerateOptions(opts ...GenerateOption) GenerateOptions {
	var options GenerateOptions
	for _, o := range opts {
		o(&options)
	}
	//set default Expiry of token
	if options.Expiry == 0 {
		options.Expiry = time.Minute * 15
	}
	return options
}
