package services

import (
	"database/sql"
	"fmt"

	"github.com/Swejal08/go-ggqlen/graph/model"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/Swejal08/go-ggqlen/utils"
	"github.com/doug-martin/goqu/v9"
)

var expenseFieldMapper = map[string]string{
	"EventID":     "event_id",
	"ItemName":    "item_name",
	"Cost":        "cost",
	"Description": "description",
	"CategoryId":  "category_id",
}

func CreateExpense(body model.NewExpense) (*model.Expense, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Insert("expense").
		Cols("event_id", "item_name", "cost", "description", "category_id").
		Vals(goqu.Vals{body.EventID, body.ItemName, body.Cost, body.Description, body.CategoryID})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	newEvent := &model.Expense{
		ID:          "1",
		EventID:     body.EventID,
		ItemName:    body.ItemName,
		Cost:        body.Cost,
		Description: body.Description,
		CategoryID:  body.CategoryID,
	}

	return newEvent, nil

}

func GetExpense(expenseId int) (*model.Expense, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "event_id", "item_name", "cost", "description", "category_id").From("expense").Where(goqu.Ex{"id": expenseId}).ToSQL()

	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	expense := &model.Expense{}

	if err := row.Scan(&expense.ID, &expense.EventID, &expense.ItemName, &expense.Cost, &expense.Description, &expense.CategoryID); err == nil {
		return expense, nil
	} else if err == sql.ErrNoRows {
		fmt.Println("No expense found", err.Error())
		return nil, err
	} else {
		fmt.Println("An error occurred while scanning row", err.Error())
		return nil, err
	}

}

func UpdateExpense(body model.UpdateExpense) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	record := utils.ConvertInputFieldsToRecord(body, expenseFieldMapper)

	ds := queryBuilder.Update("expense").Set(record).Where(goqu.Ex{"id": body.ID})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return err
	}

	if _, err = database.Exec(sql); err != nil {
		return err
	}

	return nil

}

func DeleteExpense(expenseId int) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("expense").Where(goqu.Ex{"id": expenseId})

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
