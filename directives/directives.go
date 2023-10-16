package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func CheckUserId() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		userId := ctx.Value("userId")

		if userId == nil {
			panic("Authentication Error")
		}

		return next(ctx)
	}
}
