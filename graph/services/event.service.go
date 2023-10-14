package services

import (
	"fmt"

	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/doug-martin/goqu/v9"
)

func CreateEvent(body model.NewEvent) (*model.Event, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Insert("event").
		Cols("name", "description", "location", "start_date", "end_date").
		Vals(goqu.Vals{body.Name, body.Description, body.Location, body.StartDate, body.EndDate})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	newEvent := &model.Event{
		ID:          "1",
		Name:        body.Name,
		Description: body.Description,
		Location:    body.Location,
		StartDate:   body.StartDate,
		EndDate:     body.EndDate,
	}

	return newEvent, nil

}
