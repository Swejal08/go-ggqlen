package services

import (
	"fmt"

	"github.com/Swejal08/go-ggqlen/enums"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/doug-martin/goqu/v9"
)

func CreateEventMembership(eventId int, userId int, role enums.EventMembershipRole) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

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
