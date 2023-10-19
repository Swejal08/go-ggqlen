package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Swejal08/go-ggqlen/graph/services"
)

func Authenticate() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {

		currentUserId := ctx.Value("currentUserId")

		if currentUserId == nil {
			return nil, fmt.Errorf("Authentication Failed")
		}

		userId, ok := currentUserId.(string)

		if !ok {
			return nil, fmt.Errorf("Authentication Failed")
		}

		user, err := services.GetUserById(string(userId))

		if err != nil {
			return nil, err
		}

		if user.ID == "" {
			return nil, fmt.Errorf("Authentication Failed")
		}

		return next(ctx)
	}
}
