package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/Swejal08/go-ggqlen/db"
	"github.com/Swejal08/go-ggqlen/graph/model"
	goqu "github.com/doug-martin/goqu/v9"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	database := db.GetDB()

	queryBuilder := db.GetQuilderBuilder()

	ds := queryBuilder.Insert("user").
		Cols("name", "email", "phone").
		Vals(goqu.Vals{input.Name, input.Email, input.Phone})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	if err != nil {
		fmt.Println("An error occurred while retrieving the last insert ID", err.Error())
		return nil, err
	}

	newEvent := &model.User{
		ID:    "1",
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	return newEvent, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}
