package services

import (
	"database/sql"
	"fmt"

	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func CreateEventMembership(eventId string, userId string, role model.Role) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	newId := uuid.New()

	ds := queryBuilder.Insert("event_membership").
		Cols("id", "event_id", "user_id", "role").
		Vals(goqu.Vals{newId, eventId, userId, role})

	sql, _, err := ds.ToSQL()
	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL : ", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL : ", err.Error())
	}

	return nil

}

func GetEventMembership(eventId string, userId string) (*model.EventMembership, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "event_id", "user_id", "role").From("event_membership").Where(goqu.Ex{"event_id": eventId, "user_id": userId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	membership := &model.EventMembership{}
	if err := row.Scan(&membership.ID, &membership.EventID, &membership.UserID, &membership.Role); err == nil {
		return membership, nil
	} else if err == sql.ErrNoRows {
		return nil, nil

	} else {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())
	}

}

func UpdateEventMembership(input model.AssignEventMembership, eventMembership *model.EventMembership) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Update("event_membership").Set(
		goqu.Record{"role": input.Role},
	).Where(goqu.Ex{"id": (*eventMembership).ID})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	return nil
}

func RemoveEventMembership(input model.RemoveEventMembership) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("event_membership").Where(goqu.Ex{"event_id": input.EventID, "user_id": input.UserID})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	return nil

}

func GetEventMembersDetails(eventId string) ([]*model.EventMembersDetail, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select(goqu.I("event_membership.id").As("event_membership_id"), "role", goqu.I("user.id").As("user_id"), "name", "email", "phone").From("event_membership").
		InnerJoin(goqu.T("user"), goqu.On(goqu.Ex{"event_membership.user_id": goqu.I("user.id")})).
		Where(goqu.Ex{"event_membership.event_id": eventId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	rows, err := database.Query(sqlQuery)

	if err != nil {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	defer rows.Close()

	var memberDetails []*model.EventMembersDetail

	for rows.Next() {
		eventMembers := &model.EventMembership{}
		user := &model.User{}

		if err := rows.Scan(&eventMembers.ID, &eventMembers.Role, &user.ID, &user.Name, &user.Email, &user.Phone); err != nil {

			return nil, fmt.Errorf("An error occurred while scanning rows", err.Error())
		}

		members := &model.EventMembersDetail{
			ID:   eventMembers.ID,
			Role: eventMembers.Role,
			User: user,
		}

		memberDetails = append(memberDetails, members)

		if err := rows.Err(); err != nil {

			return nil, fmt.Errorf("An error occurred after iterating through rows", err.Error())
		}
	}

	return memberDetails, nil

}
