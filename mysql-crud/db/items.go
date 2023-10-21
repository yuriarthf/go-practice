package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Item struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       float32        `json:"price"`
}

func (c *Item) Fields() []interface{} {
	return []interface{}{
		&c.ID,
		&c.Name,
		&c.Description,
		&c.Price,
	}
}

func explode(a []int64) []interface{} {
	var r []interface{}
	for i := range a {
		r = append(r, i)
	}
	return r
}

func GetAllItems() ([]Item, error) {
	r, err := db.Query("SELECT * FROM Items")
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var result []Item
	for r.Next() {
		var i Item
		if err := r.Scan(i.Fields()...); err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

func GetItemByName(name string) (Customer, error) {
	r := db.QueryRow("SELECT * FROM Items WHERE NAME = ?", name)

	var result Customer
	if err := r.Scan(result.Fields()...); err != nil {
		return result, err
	}

	return result, nil
}

func AddItem(name string, description string, price float32) (int64, error) {
	var r sql.Result
	var err error

	_, err = GetItemByName(name)
	if err == nil {
		return -1, errors.New(fmt.Sprintf(`Item with name "%s" already exist`, name))
	}

	if len(description) == 0 {
		r, err = db.Exec(
			`INSERT INTO Customers (NAME, PRICE) VALUES (?, ?)`,
			name, price,
		)
	} else {
		r, err = db.Exec(
			`INSERT INTO Customers (NAME, DESCRIPTION, PRICE) VALUES (?, ?, ?)`,
			name, description, price,
		)
	}

	if err != nil {
		return -1, err
	}

	return r.LastInsertId()
}

func itemsExist(ids []int64) (bool, error) {
	if len(ids) == 0 {
		return false, errors.New("No Item ID provided")
	}
	t := "(?" + strings.Repeat(",?", len(ids)-1) + ")"

	var count int
	db.QueryRow(
		fmt.Sprintf("SELECT COUNT(*) FROM Items WHERE ID IN %s", t),
		explode(ids)...,
	).Scan(&count)

	if count == len(ids) {
		return true, nil
	}

	return false, nil
}
