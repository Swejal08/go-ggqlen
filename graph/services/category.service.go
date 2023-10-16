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

var categoryFieldMapper = map[string]string{
	"CategoryName": "category_name",
}

func CreateCategory(body model.NewCategory) (*model.Category, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	newId := uuid.New()

	ds := queryBuilder.Insert("category").
		Cols("id", "category_name").
		Vals(goqu.Vals{newId, body.CategoryName})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	newCategory := &model.Category{
		ID:           newId.String(),
		CategoryName: body.CategoryName,
	}

	return newCategory, nil

}

func GetCategory(categoryId string) (*model.Category, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "category_name").From("category").Where(goqu.Ex{"id": categoryId}).ToSQL()

	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	category := &model.Category{}

	if err := row.Scan(&category.ID, &category.CategoryName); err == nil {
		return category, nil
	} else if err == sql.ErrNoRows {
		fmt.Println("No category found", err.Error())
		return nil, err
	} else {
		fmt.Println("An error occurred while scanning row", err.Error())
		return nil, err
	}

}

func UpdateCategory(body model.UpdateCategory) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	record := utils.ConvertInputFieldsToRecord(body, categoryFieldMapper)

	ds := queryBuilder.Update("category").Set(record).Where(goqu.Ex{"id": body.ID})

	sql, _, err := ds.ToSQL()

	if err != nil {
		return err
	}

	if _, err = database.Exec(sql); err != nil {
		return err
	}

	return nil

}

func DeleteCategory(categoryId string) error {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	ds := queryBuilder.Delete("category").Where(goqu.Ex{"id": categoryId})

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
