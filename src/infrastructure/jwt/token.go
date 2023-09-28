package jwt

import (
	"context"
	"errors"
	"github.com/iAlijon/eater-service/src/infrastructure/rand"

)

type Token struct {
	UserID string
}

type ctxKey int

const (
	ctxTokenKey ctxKey = iota
)

func TokenFromCtx(ctx context.Context) (*Token, error) {
	info, ok := ctx.Value(ctxTokenKey).(*Token)
	if ok {
		return info, nil
	}
	return nil, errors.New("token not found")
}

func TokenInCtx(parent context.Context, token *Token) context.Context {
	return context.WithValue(parent, ctxTokenKey, token)
}
