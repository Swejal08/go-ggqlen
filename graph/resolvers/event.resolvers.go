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

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	if err := utils.ValidateInput(input); err != nil {
		return nil, err
	}

	userId := ctx.Value("userId").(string)

	event, err := services.CreateEvent(input)

	if err != nil {
		return nil, err
	}

	err = services.CreateEventMembership(event.ID, userId, "admin")

	if err != nil {
		return nil, err
	}

	return event, nil
}

// UpdateEvent is the resolver for the updateEvent field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, input model.UpdateEvent) (*string, error) {
	if err := utils.ValidateInput(input); err != nil {
		return nil, err
	}

	userId := ctx.Value("userId").(string)

	allowedRoles := []enums.EventMembershipRole{enums.Admin, enums.Contributor}

	err := accessControl.Check(allowedRoles, userId, input.ID)

	if err != nil {
		return nil, fmt.Errorf("This resource is forbidden")
	}

	_, err = services.GetEvent(input.ID)

	if err != nil {
		return nil, err
	}

	err = services.UpdateEvent(input)

	if err != nil {
		return nil, err
	}

	successMessage := "Event has been updated"
	return &successMessage, nil
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, input model.DeleteEvent) (*string, error) {
	userId := ctx.Value("userId").(string)

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	accessError := accessControl.Check(allowedRoles, userId, input.ID)

	if accessError != nil {
		return nil, accessError
	}

	_, err := services.GetEvent(input.ID)

	if err != nil {
		return nil, err
	}

	/*
		Todo: Need to soft delete
		goqu does not have softDelete out of the box
	*/

	err = services.DeleteEvent(input.ID)

	if err != nil {
		return nil, err
	}

	successMessage := "Event has been deleted"
	return &successMessage, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	//TODO apply pagination

	userId := ctx.Value("userId").(string)

	events, err := services.GetMyEvents(userId)

	if err != nil {
		return nil, err
	}

	return events, nil
}

// EventDetails is the resolver for the eventDetails field.
func (r *queryResolver) EventDetails(ctx context.Context, eventID string) (*model.EventDetails, error) {
	userId := ctx.Value("userId").(string)

	allowedRoles := []enums.EventMembershipRole{enums.Admin, enums.Contributor, enums.Attendee}

	accessError := accessControl.Check(allowedRoles, userId, eventID)

	if accessError != nil {
		return nil, accessError
	}

	event, err := services.GetEvent(eventID)

	sessions, err := services.GetSessionByEventId(eventID)

	if err != nil {
		return nil, err
	}

	eventDetails := &model.EventDetails{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		Sessions:    sessions,
	}

	return eventDetails, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
