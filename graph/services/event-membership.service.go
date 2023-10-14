package services

import (
	"database/sql"
	"fmt"

	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/doug-martin/goqu/v9"
)

func CreateEventMembership(eventId int, userId int, role model.Role) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	// there could be more appropriate way to use enums rather than this

	// membershipRole := enums.GetRoleDescription(int(role))

	ds := queryBuilder.Insert("event_membership").
		Cols("event_id", "user_id", "role").
		Vals(goqu.Vals{eventId, userId, role})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return err
	}

	fmt.Println("Event Membership has been created")

	return nil

}

func GetEventMembership(eventId int, userId int) *model.EventMembership {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "event_id", "user_id", "role").From("event_membership").Where(goqu.Ex{"event_id": eventId, "user_id": userId}).ToSQL()

	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	membership := &model.EventMembership{}
	if err := row.Scan(&membership.ID, &membership.EventID, &membership.UserID, &membership.Role); err == nil {
		return membership
	} else if err == sql.ErrNoRows {
		return nil
	} else {
		fmt.Println("An error occurred while scanning row", err.Error())
		return nil
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
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return err
	}

	return nil
}

func RemoveEventMembership(input model.RemoveEventMembership) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("event_membership").Where(goqu.Ex{"event_id": input.EventID, "user_id": input.UserID})

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
