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

	newId := uuid.New()

	ds := queryBuilder.Insert("expense").
		Cols("id", "event_id", "item_name", "cost", "description", "category_id").
		Vals(goqu.Vals{newId, body.EventID, body.ItemName, body.Cost, body.Description, body.CategoryID})

	sql, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	newEvent := &model.Expense{
		ID:          newId.String(),
		EventID:     body.EventID,
		ItemName:    body.ItemName,
		Cost:        body.Cost,
		Description: body.Description,
	}

	return newEvent, nil

}

func GetExpense(expenseId string) (*model.Expense, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "event_id", "item_name", "cost", "description", "category_id").From("expense").Where(goqu.Ex{"id": expenseId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	expense := &model.Expense{}

	if err := row.Scan(&expense.ID, &expense.EventID, &expense.ItemName, &expense.Cost, &expense.Description, &expense.CategoryID); err == nil {
		return expense, nil
	} else if err == sql.ErrNoRows {
		return nil, fmt.Errorf("No expense found", err.Error())

	} else {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

}

func UpdateExpense(body model.UpdateExpense) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	record := utils.ConvertInputFieldsToRecord(body, expenseFieldMapper)

	ds := queryBuilder.Update("expense").Set(record).Where(goqu.Ex{"id": body.ID})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL", err.Error())

	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	return nil

}

func DeleteExpense(expenseId string) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("expense").Where(goqu.Ex{"id": expenseId})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return fmt.Errorf("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		return fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	return nil

}

func GetTotalExpensesBasedOnCategory(event *model.Event) (*model.TotalExpense, error) {

	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.
		Select(
			goqu.I("category.id").As("category_id"),
			goqu.I("category.category_name").As("category_name"),
			goqu.SUM("expense.cost").As("expense"),
		).
		From("event").InnerJoin(goqu.T("expense"), goqu.On(goqu.Ex{"event.id": goqu.I("event_id")})).
		InnerJoin(goqu.T("category"), goqu.On(goqu.Ex{"category_id": goqu.I("category.id")})).Where(goqu.Ex{"event.id": event.ID}).GroupBy("category.id", "category.category_name")

	sqlQuery, _, err := ds.ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL", err.Error())

	}

	rows, err := database.Query(sqlQuery)

	if err != nil {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	defer rows.Close()

	var categoryResponse []*model.CategoryExpense
	totalExpense := 0
	for rows.Next() {
		category := &model.CategoryExpense{}
		if err := rows.Scan(&category.ID, &category.Name, &category.Expense); err != nil {
			return nil, fmt.Errorf("An error occurred while scanning rows", err.Error())
		}

		categoryResponse = append(categoryResponse, category)
		totalExpense += *&category.Expense

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("An error occurred after iterating through rows", err.Error())
	}

	expense := &model.TotalExpense{
		TotalExpense: totalExpense,
		Name:         event.Name,
		Category:     categoryResponse,
	}

	return expense, nil

}