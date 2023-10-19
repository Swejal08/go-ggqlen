package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	accessControl "github.com/Swejal08/go-ggqlen/access-control"
	"github.com/Swejal08/go-ggqlen/enums"
	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/graph/services"
	"github.com/Swejal08/go-ggqlen/utils"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	if err := utils.ValidateInput(input); err != nil {
		return nil, err
	}

	userId := ctx.Value("currentUserId").(string)

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	accessError := accessControl.Check(allowedRoles, userId, input.EventID)

	if accessError != nil {
		return nil, accessError
	}

	category, err := services.CreateCategory(input)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, input model.UpdateCategory) (*string, error) {
	if err := utils.ValidateInput(input); err != nil {
		return nil, err
	}
	userId := ctx.Value("currentUserId").(string)

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	accessError := accessControl.Check(allowedRoles, userId, input.EventID)

	if accessError != nil {
		return nil, accessError
	}

	category, err := services.GetCategory(input.ID)

	if category == nil {
		return nil, err
	}

	err = services.UpdateCategory(input)

	if err != nil {
		return nil, err
	}

	successMessage := "Category has been updated"
	return &successMessage, nil
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, input model.DeleteCategory) (*string, error) {
	userId := ctx.Value("currentUserId").(string)

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	accessError := accessControl.Check(allowedRoles, userId, input.EventID)

	if accessError != nil {
		return nil, accessError
	}

	category, err := services.GetCategory(input.ID)

	if category == nil {
		return nil, err
	}

	err = services.DeleteCategory(input.ID)

	if err != nil {
		return nil, err

	}

	successMessage := "Category has been deleted"
	return &successMessage, nil
}

// GetCategories is the resolver for the getCategories field.
func (r *queryResolver) GetCategories(ctx context.Context) ([]*model.Category, error) {
	// userId := ctx.Value("currentUserId").(string)

	// allowedRoles := []enums.EventMembershipRole{enums.Admin}

	// accessError := accessControl.Check(allowedRoles, userId, eventID)

	// if accessError != nil {
	// 	return nil, accessError
	// }

	categories, err := services.GetAllCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
