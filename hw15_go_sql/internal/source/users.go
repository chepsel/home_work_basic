package source

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserStat struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	OrderSum float32 `json:"orderSum,omitempty"`
	AvgPrice float32 `json:"avgPrice,omitempty"`
}

func (ctr *User) MarshalJSON() ([]byte, error) {
	type dropDefaultInf User
	result, err := json.Marshal((*dropDefaultInf)(ctr))
	return result, err
}

func (ctr *User) UnmarshalJSON(data []byte) error {
	type dropDefaultInf User
	err := json.Unmarshal(data, (*dropDefaultInf)(ctr))
	if err != nil {
		return err
	}
	return nil
}

func (ctr *UserStat) MarshalJSON() ([]byte, error) {
	type dropDefaultInf UserStat
	result, err := json.Marshal((*dropDefaultInf)(ctr))
	return result, err
}

func (ctr *UserStat) UnmarshalJSON(data []byte) error {
	type dropDefaultInf UserStat
	err := json.Unmarshal(data, (*dropDefaultInf)(ctr))
	if err != nil {
		return err
	}
	return nil
}

func (src *Database) UnmarshalJSONUsersStat(data []byte) ([]UserStat, error) {
	var ctr []UserStat
	if err := json.Unmarshal(data, &ctr); err != nil {
		return nil, err
	}
	return ctr, nil
}

func (src *Database) MarshalJSONUsersStat(ctr []UserStat) ([]byte, error) {
	result, err := json.Marshal(ctr)
	return result, err
}

func (src *Database) UnmarshalJSONUsers(data []byte) ([]User, error) {
	var ctr []User
	if err := json.Unmarshal(data, &ctr); err != nil {
		return nil, err
	}
	return ctr, nil
}

func (src *Database) MarshalJSONUsers(ctr []User) ([]byte, error) {
	result, err := json.Marshal(ctr)
	return result, err
}

func (src *Database) NewUser() *User {
	return &User{}
}

func (src *Database) NewUserStat() *UserStat {
	return &UserStat{}
}

func (src *Database) InsertUser(user *User) error {
	queryType := "InsertUser"
	query := `INSERT INTO users (name, email, password) VALUES($1, $2, $3) RETURNING id`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%s,%s,%s", user.Name, user.Email, user.Password)) // Log

	err := src.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return src.LogError(queryType, err)
	}
	go src.LogDBResult(queryType, fmt.Sprintf("%d", user.ID)) // Log

	if user.ID == 0 {
		return src.LogError(queryType, MissingID)
	}
	return nil
}

func (src *Database) DeleteUser(id int64) error {
	queryType := "DeleteUser"
	query := `DELETE FROM users WHERE id=$1`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", id)) // Log

	result, err := src.DB.Exec(query, id)
	if err != nil {
		return src.LogError(queryType, err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		src.LogError(queryType, NoRowsAffected)
		return NoRowsAffected
	}

	go src.LogDBResult(queryType, fmt.Sprintf("deleted rows(%d)", rowsAffected)) // Log

	return nil
}

func (src *Database) UpdateUser(user *User) error {
	queryType := "UpdateUser"
	query := user.GetUpdateQuery()
	go src.LogDBRequest(queryType, query, fmt.Sprintf("%s,%s,%s - where id='%d'",
		user.Name,
		user.Email,
		user.Password,
		user.ID,
	)) // Log

	result, err := src.DB.Exec(query, user.ID)
	if err != nil {
		return src.LogError(queryType, err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		src.LogError(queryType, NoRowsAffected)
		return NoRowsAffected
	}

	go src.LogDBResult(queryType, fmt.Sprintf("updated rows(%d)", rowsAffected)) // Log

	return nil
}

func (ctr *User) GetUpdateQuery() string {
	var query string
	switch {
	case len(ctr.Name) > 0 && len(ctr.Email) > 0 && len(ctr.Password) > 0:
		query = fmt.Sprintf(`UPDATE ctrs SET name='%s', email='%s', password='%s' WHERE ID=$1;`,
			ctr.Name,
			ctr.Email,
			ctr.Password,
		)
	case len(ctr.Name) > 0 && len(ctr.Email) > 0 && len(ctr.Password) == 0:
		query = fmt.Sprintf(`UPDATE ctrs SET name='%s', email='%s' WHERE ID=$1;`, ctr.Name, ctr.Email)
	case len(ctr.Name) > 0 && len(ctr.Email) == 0 && len(ctr.Password) > 0:
		query = fmt.Sprintf(`UPDATE ctrs SET name='%s', password='%s' WHERE ID=$1;`, ctr.Name, ctr.Password)
	case len(ctr.Name) == 0 && len(ctr.Email) > 0 && len(ctr.Password) > 0:
		query = fmt.Sprintf(`UPDATE ctrs SET email='%s', password='%s' WHERE ID=$1;`, ctr.Email, ctr.Password)
	case len(ctr.Name) > 0 && len(ctr.Email) == 0 && len(ctr.Password) == 0:
		query = fmt.Sprintf(`UPDATE ctrs SET name='%s' WHERE ID=$1;`, ctr.Name)
	case len(ctr.Name) == 0 && len(ctr.Email) > 0 && len(ctr.Password) == 0:
		query = fmt.Sprintf(`UPDATE ctrs SET email='%s' WHERE ID=$1;`, ctr.Email)
	case len(ctr.Name) == 0 && len(ctr.Email) == 0 && len(ctr.Password) > 0:
		query = fmt.Sprintf(`UPDATE ctrs SET password='%s' WHERE ID=$1;`, ctr.Password)
	}
	return query
}

func (src *Database) GetUser(userID int64) (User, error) {
	queryType := "GetUser"
	var user User

	query := `SELECT id, name, email, password FROM users WHERE id=$1`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", userID)) // Log

	err := src.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, src.LogError(queryType, err)
	}
	go src.LogDBResult(queryType, fmt.Sprintf("%v,%v,%v,%v",
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)) // Log

	return user, nil
}

func (src *Database) GetUsersList() ([]User, error) {
	queryType := "GetUsersList"
	var users []User

	query := `SELECT id, name, email, password FROM users`

	go src.LogDBRequest(queryType, query, "") // Log

	rows, err := src.DB.Query(query)
	if err != nil {
		return nil, src.LogError(queryType, err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err = user.Scan(rows); err != nil {
			return nil, src.LogError(queryType, err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, src.LogError(queryType, err)
	}

	go src.LogDBResult(queryType, fmt.Sprintf("rows found(%d)", len(users))) // Log

	return users, nil
}

func (src *Database) GetUserStat(userID int64) (UserStat, error) {
	queryType := "GetUserStat"
	var userStat UserStat

	query := `SELECT u.id, u."name", SUM(o.total_amount) AS order_sum , AVG(p.price) AS avg_price 
	FROM (SELECT * FROM orders where user_id=$1) o 
	INNER JOIN orderproducts op ON op.order_id=o.id
	INNER JOIN products p ON op.product_id=p.id
	INNER JOIN users u ON o.user_id=u.id
	GROUP BY u."name", u.id`

	go src.LogDBRequest(queryType, query, fmt.Sprintf("%d", userID)) // Log

	err := src.DB.QueryRow(query, userID).Scan(&userStat.ID, &userStat.Name, &userStat.OrderSum, &userStat.AvgPrice)
	if err != nil {
		return userStat, src.LogError(queryType, err)
	}
	go src.LogDBResult(queryType, fmt.Sprintf("%v,%v,%v,%v",
		userStat.ID,
		userStat.Name,
		userStat.AvgPrice,
		userStat.OrderSum,
	)) // Log

	return userStat, nil
}

func (src *Database) GetUsersStat() ([]UserStat, error) {
	queryType := "GetUsersStat"
	var usersStat []UserStat

	query := `SELECT u.id, u."name", SUM(o.total_amount) AS order_sum , AVG(p.price) AS avg_price
	FROM orders o 
	INNER JOIN orderproducts op ON op.order_id=o.id
	INNER JOIN products p ON op.product_id=p.id
	INNER JOIN users u ON o.user_id=u.id
	GROUP BY u."name", u.id`

	go src.LogDBRequest(queryType, query, "") // Log

	rows, err := src.DB.Query(query)
	if err != nil {
		return nil, src.LogError(queryType, err)
	}

	defer rows.Close()

	for rows.Next() {
		var userStat UserStat
		if err = userStat.Scan(rows); err != nil {
			return nil, src.LogError(queryType, err)
		}
		usersStat = append(usersStat, userStat)
	}

	if err = rows.Err(); err != nil {
		return nil, src.LogError(queryType, err)
	}

	go src.LogDBResult(queryType, fmt.Sprintf("rows found(%d)", len(usersStat))) // Log

	return usersStat, nil
}

func (ctr *User) Scan(rows *sql.Rows) error {
	if err := rows.Scan(&ctr.ID, &ctr.Name, &ctr.Email, &ctr.Password); err != nil {
		return err
	}
	return nil
}

func (ctr *UserStat) Scan(rows *sql.Rows) error {
	if err := rows.Scan(&ctr.ID, &ctr.Name, &ctr.OrderSum, &ctr.AvgPrice); err != nil {
		return err
	}
	return nil
}
