package services

import (
	"fmt"

	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/doug-martin/goqu/v9"
)

func CreateUser(body model.NewUser) (*model.User, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Insert("user").
		Cols("name", "email", "phone").
		Vals(goqu.Vals{body.Name, body.Email, body.Phone})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	newEvent := &model.User{
		ID:    "1",
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}

	return newEvent, nil
}
