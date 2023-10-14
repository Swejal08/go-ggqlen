package services

import (
	"database/sql"
	"fmt"

	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/Swejal08/go-ggqlen/utils"
	"github.com/doug-martin/goqu/v9"
)

var fieldMapper = map[string]string{
	"Name":        "name",
	"Description": "description",
	"Location":    "location",
	"StartDate":   "start_date",
	"EndDate":     "end_date",
}

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

func GetEvent(eventId int) (*model.Event, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "name", "description", "location", "start_date", "end_date").From("event").Where(goqu.Ex{"id": eventId}).ToSQL()

	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	event := &model.Event{}

	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.StartDate, &event.EndDate); err == nil {
		return event, nil
	} else if err == sql.ErrNoRows {
		fmt.Println("No events found", err.Error())
		return nil, err
	} else {
		fmt.Println("An error occurred while scanning row", err.Error())
		return nil, err
	}

}

func UpdateEvent(body model.UpdateEvent) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	record := utils.ConvertInputFieldsToRecord(body, fieldMapper)

	ds := queryBuilder.Update("event").Set(record).Where(goqu.Ex{"id": body.ID})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return err
	}

	if _, err = database.Exec(sql); err != nil {
		return err
	}

	return nil

}

func DeleteEvent(eventId int) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("event").Where(goqu.Ex{"id": eventId})

	sql, _, err := ds.ToSQL()

	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return err
	}

	return nil

}
