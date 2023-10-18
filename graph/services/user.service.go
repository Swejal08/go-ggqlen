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

func CreateUser(body model.NewUser) (*model.User, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	newId := uuid.New()

	hashedPassword, err := utils.HashPassword((body.Password))

	if err != nil {
		return nil, err
	}

	ds := queryBuilder.Insert("user").
		Cols("id", "name", "email", "phone", "password").
		Vals(goqu.Vals{newId, body.Name, body.Email, body.Phone, hashedPassword})

	sql, _, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	}

	if _, err = database.Exec(sql); err != nil {
		fmt.Println("An error occurred while executing the SQL", err.Error())
		return nil, err
	}

	user := &model.User{
		ID:    newId.String(),
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}

	return user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "name", "email", "phone", "password").
		From("user").
		Where(goqu.Ex{"email": email}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL: ", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	user := &model.User{}

	row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password)

	return user, err

}

func GetUserById(userId string) (*model.User, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select("id", "name", "email", "phone").
		From("user").
		Where(goqu.Ex{"id": userId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL: ", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	user := &model.User{}

	row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)

	return user, err
}

func GetUserDetailsForEvent(memberId string, eventId string) (*model.UserDetails, error) {
	database := initializer.GetDB()

	queryBuilder := initializer.GetQueryBuilder()

	sqlQuery, _, err := queryBuilder.Select(goqu.I("user.id").As("id"), "name", "email", "phone", "role").
		From("user").
		InnerJoin(goqu.T("event_membership"), goqu.On(goqu.Ex{"user.id": goqu.I("event_membership.user_id")})).
		Where(
			goqu.And(
				goqu.Ex{"event_membership.event_id": eventId},
				goqu.Ex{"user.id": memberId},
			),
		).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("An error occurred while generating the SQL: ", err.Error())
	}

	row := database.QueryRow(sqlQuery)

	userDetail := &model.UserDetails{}

	if err := row.Scan(&userDetail.ID, &userDetail.Name, &userDetail.Email, &userDetail.Phone, &userDetail.Role); err == nil {

		return userDetail, nil
	} else if err == sql.ErrNoRows {
		return nil, fmt.Errorf("User does not have membership: ", err.Error())
	} else {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())
	}
}

func GetNonEventMembers(eventId string) ([]*model.User, error) {
	database := initializer.GetDB()

	var sql string

	sql = `
    SELECT u.id AS user_id, u.name, u.email, u.phone
    FROM "user" u
    LEFT JOIN event_membership em ON em.user_id = u.id
    WHERE u.id NOT IN (
        SELECT em."user_id"
        FROM event_membership em
        WHERE em.event_id = $1
    ) AND em.event_id IS NULL
`

	rows, err := database.Query(sql, eventId)

	if err != nil {
		return nil, fmt.Errorf("An error occurred while executing the SQL", err.Error())

	}

	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone); err != nil {

			return nil, fmt.Errorf("An error occurred while scanning rows", err.Error())
		}

		users = append(users, user)

	}

	if err := rows.Err(); err != nil {

		return nil, fmt.Errorf("An error occurred after iterating through rows", err.Error())
	}

	return users, nil
}
