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

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input model.NewExpense) (*model.Expense, error) {
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

	event, err := services.CreateExpense(input)

	if err != nil {
		fmt.Println("Expense cannot be created", err.Error())
	}

	return event, nil
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, input model.UpdateExpense) (*string, error) {
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

	expense, err := services.GetExpense(id)

	if expense == nil {
		return nil, err
	}

	err = services.UpdateExpense(input)

	if err != nil {
		fmt.Println("Something went wrong when updating expense", err.Error())
	}

	successMessage := "Expense has been updated"
	return &successMessage, nil
}

// DeleteExpense is the resolver for the deleteExpense field.
func (r *mutationResolver) DeleteExpense(ctx context.Context, input model.DeleteExpense) (*string, error) {
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

	expense, err := services.GetExpense(id)

	fmt.Println(expense)

	if expense == nil {
		return nil, err
	}

	err = services.DeleteExpense(id)

	if err != nil {
		fmt.Println("Something went wrong when deleting expense", err.Error())
	}

	successMessage := "Expense has been deleted"
	return &successMessage, nil
}
