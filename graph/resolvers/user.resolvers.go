package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	accessControl "github.com/Swejal08/go-ggqlen/access-control"
	"github.com/Swejal08/go-ggqlen/enums"
	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/graph/services"
	"github.com/Swejal08/go-ggqlen/utils"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if err := utils.ValidateInput(input); err != nil {
		return nil, err
	}

	user, err := services.GetUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if user.ID != "" {
		return nil, fmt.Errorf("User already exists")
	}

	createdUser, err := services.CreateUser(input)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// MyUserDetail is the resolver for the myUserDetail field.
func (r *queryResolver) MyUserDetail(ctx context.Context, eventID string) (*model.UserDetails, error) {
	userId := ctx.Value("currentUserId").(string)

	user, _ := services.GetUserDetailsForEvent(userId, eventID)

	if user.ID == "" {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

// NonEventMembers is the resolver for the nonEventMembers field.
func (r *queryResolver) NonEventMembers(ctx context.Context, eventID string) ([]*model.User, error) {
	userId := ctx.Value("currentUserId").(string)

	allowedRoles := []enums.EventMembershipRole{enums.Admin, enums.Contributor}

	accessError := accessControl.Check(allowedRoles, userId, eventID)

	if accessError != nil {
		return nil, accessError
	}

	users, err := services.GetNonEventMembers(eventID)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// UserDetails is the resolver for the userDetails field.
func (r *queryResolver) UserDetails(ctx context.Context, userID string, eventID string) (*model.UserDetails, error) {
	userId := ctx.Value("currentUserId").(string)
	allowedRoles := []enums.EventMembershipRole{enums.Admin, enums.Contributor}

	accessError := accessControl.Check(allowedRoles, userId, eventID)

	if accessError != nil {
		return nil, accessError
	}

	userDetails, err := services.GetUserDetailsForEvent(userID, eventID)

	if err != nil {
		return nil, err
	}

	return userDetails, nil
}
