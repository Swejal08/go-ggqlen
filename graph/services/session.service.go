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

var sessionFieldMapper = map[string]string{
	"EventID":   "event_id",
	"StartDate": "start_date",
	"EndDate":   "end_date",
}

func CreateSession(body model.NewSession) (*model.Session, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	newId := uuid.New()

	ds := queryBuilder.Insert("sessions").
		Cols("id", "event_id", "start_date", "end_date").
		Vals(goqu.Vals{newId, body.EventID, body.StartDate, body.EndDate})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	newSession := &model.Session{
		ID:        newId.String(),
		EventID:   body.EventID,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	}

	return newSession, nil
}

func GetSession(sessionId string) (*model.Session, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "event_id", "start_date", "end_date").From("sessions").Where(goqu.Ex{"id": sessionId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	session := &model.Session{}

	if err := row.Scan(&session.ID, &session.EventID, &session.StartDate, &session.EndDate); err == nil {
		return session, nil
	} else if err == sql.ErrNoRows {

		return nil, fmt.Errorf("No session found", err.Error())
	} else {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

}

func GetSessionByEventId(eventId string) ([]*model.Session, error) {
	fmt.Println(eventId)
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "event_id", "start_date", "end_date").From("sessions").Where(goqu.Ex{"event_id": eventId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	rows, err := database.Query(sqlQuery)

	if err != nil {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	defer rows.Close()

	var sessions []*model.Session

	for rows.Next() {
		session := &model.Session{}
		if err := rows.Scan(&session.ID, &session.EventID, &session.StartDate, &session.EndDate); err != nil {
			return nil, fmt.Errorf("An error occurred while scanning rows", err.Error())

		}

		sessions = append(sessions, session)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("An error occurred after iterating through rows", err.Error())

	}

	return sessions, nil
}

func UpdateSession(body model.UpdateSession) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	record := utils.ConvertInputFieldsToRecord(body, sessionFieldMapper)

	ds := queryBuilder.Update("sessions").Set(record).Where(goqu.Ex{"id": body.ID})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL", err.Error())

	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	return nil

}

func DeleteSession(sessionId string) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("sessions").Where(goqu.Ex{"id": sessionId})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	return nil

}
