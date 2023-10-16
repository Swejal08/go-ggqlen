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
	"github.com/Swejal08/go-ggqlen/initializer"
	goqu "github.com/doug-martin/goqu/v9"
)

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	event, err := services.CreateEvent(input)

	if err != nil {
		fmt.Println("Event cannot be created", err.Error())
	}

	// replace 1 and 1 with  eventId that will come from event and userId from ctx.

	err = services.CreateEventMembership(1, uId, "admin")

	if err != nil {
		fmt.Println("Event Membership cannot be created", err.Error())
	}

	return event, nil
}

// UpdateEvent is the resolver for the updateEvent field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, input model.UpdateEvent) (*string, error) {
	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	eventId, err := strconv.Atoi(input.ID)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	allowedRoles := []enums.EventMembershipRole{enums.Admin, enums.Contributor}

	hasAccess := accessControl.Check(allowedRoles, uId, eventId)

	if !hasAccess {
		panic("Access denied")
	}

	//need to convert db id type to string

	id, err := strconv.Atoi(input.ID)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	event, err := services.GetEvent(id)

	if event == nil {
		return nil, err
	}

	err = services.UpdateEvent(input)

	if err != nil {
		fmt.Println("Something went wrong when updating event", err.Error())
	}

	successMessage := "Event has been updated"
	return &successMessage, nil
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, input model.DeleteEvent) (*string, error) {
	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	eventId, err := strconv.Atoi(input.ID)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	hasAccess := accessControl.Check(allowedRoles, uId, eventId)

	if !hasAccess {
		panic("Access denied")
	}

	event, err := services.GetEvent(eventId)

	if event == nil {
		return nil, err
	}

	/*
		Todo: Need to soft delete
		goqu does not have softDelete out of the box
	*/

	err = services.DeleteEvent(eventId)

	if err != nil {
		fmt.Println("Something went wrong when deleting event", err.Error())
	}

	successMessage := "Event has been deleted"
	return &successMessage, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	if err != nil {
		fmt.Println("error converting ID to int: %w", err)
	}

	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Select(
		goqu.I("event.id").As("event_id"), "name", "description",
		"location", "start_date", "end_date").
		From("event_membership").InnerJoin(goqu.T("event"), goqu.On(goqu.Ex{"event_id": goqu.I("event.id")})).Where(goqu.Ex{"event_membership.user_id": uId})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	rows, err := database.Query(sql)

	if err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	defer rows.Close()

	var events []*model.Event

	for rows.Next() {
		event := &model.Event{}
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.StartDate, &event.EndDate); err != nil {
			fmt.Println("An error occurred while scanning rows", err.Error())
			return nil, err
		}

		events = append(events, event)

	}

	if err := rows.Err(); err != nil {
		fmt.Println("An error occurred after iterating through rows", err.Error())
		return nil, err
	}

	return events, nil
}

// Event is the resolver for the event field.
func (r *queryResolver) Event(ctx context.Context, eventID int) (*model.Event, error) {
	userId := ctx.Value("userId").(string)

	uId, err := strconv.Atoi(userId)

	allowedRoles := []enums.EventMembershipRole{enums.Admin}

	hasAccess := accessControl.Check(allowedRoles, uId, eventID)

	if !hasAccess {
		panic("You do not have event membership")
	}

	event, err := services.GetEvent(eventID)

	if err != nil {
		return nil, err
	}

	return event, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
