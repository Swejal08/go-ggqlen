package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"
	"strconv"

	accessControl "github.com/Swejal08/go-ggqlen/access-control"
	"github.com/Swejal08/go-ggqlen/enums"
	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/graph/services"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {

	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	hasAccess := accessControl.Check(allowedRoles, uId, input.EventID)

	if !hasAccess {
		panic("Access denied")
	}

	category, err := services.CreateCategory(input)

	if err != nil {
		fmt.Println("Category cannot be created", err.Error())
	}

	return category, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, input model.UpdateCategory) (*string, error) {

	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	id, err := strconv.Atoi(input.ID)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	hasAccess := accessControl.Check(allowedRoles, uId, input.EventID)

	if !hasAccess {
		panic("Access denied")
	}

	category, err := services.GetCategory(id)

	if category == nil {
		return nil, err
	}

	err = services.UpdateCategory(input)

	if err != nil {
		fmt.Println("Something went wrong when updating category", err.Error())
	}

	successMessage := "Category has been updated"
	return &successMessage, nil
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, input model.DeleteCategory) (*string, error) {
	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	id, err := strconv.Atoi(input.ID)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	hasAccess := accessControl.Check(allowedRoles, uId, input.EventID)

	if !hasAccess {
		panic("Access denied")
	}
	category, err := services.GetCategory(id)

	if category == nil {
		return nil, err
	}

	err = services.DeleteCategory(id)

	if err != nil {
		fmt.Println("Something went wrong when deleting category", err.Error())
	}

	successMessage := "Category has been deleted"
	return &successMessage, nil
}
