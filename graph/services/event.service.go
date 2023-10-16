package services

import (
	"database/sql"
	"fmt"

	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/Swejal08/go-ggqlen/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

var eventFieldMapper = map[string]string{
	"Name":        "name",
	"Description": "description",
	"Location":    "location",
}

func CreateEvent(body model.NewEvent) (*model.Event, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	newId := uuid.New()

	ds := queryBuilder.Insert("event").
		Cols("id", "name", "description", "location").
		Vals(goqu.Vals{newId, body.Name, body.Description, body.Location})

	sql, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL:", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return nil, fmt.Errorf("An error occurred while executing the SQL: ", err.Error())
	}

	newEvent := &model.Event{
		ID:          newId.String(),
		Name:        body.Name,
		Description: body.Description,
		Location:    body.Location,
	}

	return newEvent, nil

}

func GetEvent(eventId string) (*model.Event, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "name", "description", "location").From("event").Where(goqu.Ex{"id": eventId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL: ", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	event := &model.Event{}

	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location); err == nil {
		return event, nil
	} else if err == sql.ErrNoRows {
		return nil, fmt.Errorf("No event found: ", err.Error())
	} else {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())
	}

}

func GetMyEvents(userId string) ([]*model.Event, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Select(
		goqu.I("event.id").As("event_id"), "name", "description",
		"location").
		From("event_membership").InnerJoin(goqu.T("event"), goqu.On(goqu.Ex{"event_id": goqu.I("event.id")})).Where(goqu.Ex{"event_membership.user_id": userId})

	sql, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	rows, err := database.Query(sql)

	if err != nil {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	defer rows.Close()

	var events []*model.Event

	for rows.Next() {
		event := &model.Event{}
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location); err != nil {

			return nil, fmt.Errorf("An error occurred while scanning rows", err.Error())
		}

		events = append(events, event)

	}

	if err := rows.Err(); err != nil {

		return nil, fmt.Errorf("An error occurred after iterating through rows", err.Error())
	}

	return events, nil
}

func UpdateEvent(body model.UpdateEvent) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	record := utils.ConvertInputFieldsToRecord(body, eventFieldMapper)

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

func DeleteEvent(eventId string) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("event").Where(goqu.Ex{"id": eventId})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL: ", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL: ", err.Error())
	}

	return nil

}
