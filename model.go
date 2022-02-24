package main

import (
	"database/sql"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *user) getUser(db *sql.DB) error {
	//statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", u.ID)
	//return db.QueryRow(statement).Scan(&u.Name, &u.Age)

	statement := "SELECT name, age FROM users WHERE id = ?"
	return db.QueryRow(statement, u.ID).Scan(&u.Name, &u.Age)
}

func (u *user) updateUser(db *sql.DB) error {
	//statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)
	//_, err := db.Exec(statement)

	statement := "update users set name = ?, age = ? where id = ?"
	_, err := db.Exec(statement, u.Name, u.Age, u.ID)

	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	//statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
	//_, err := db.Exec(statement)

	statement := "delete from users where id = ?"
	_, err := db.Exec(statement, u.ID)

	return err

}

func (u *user) createUser(db *sql.DB) error {
	//statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", u.Name, u.Age)
	//_, err := db.Exec(statement)

	statement := "insert into users(name, age) values(?, ?)"
	_, err := db.Exec(statement, u.Name, u.Age)

	if err != nil {
		return err
	}

	err = db.QueryRow("select last_insert_id()").Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	//statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)
	//rows, err := db.Query(statement)

	statement := "select id, name, age from users limit ? offset ?"
	rows, err := db.Query(statement, count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
