package data

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

func CreateClient() (db *sql.DB, err error) {
	return sql.Open("sqlite3", "./example.db")
}

func GetAllItems() (result []Item, err error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		item := Item{}

		err = rows.Scan(&item.ID, &item.Name, &item.Notes, &item.Description, &item.Stage, &item.Category)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return
}

func GetItemByID(id int) (item Item, err error) {
	db, err := CreateClient()
	if err != nil {
		return Item{}, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM items WHERE id = ?", id)
	err = row.Scan(&item.ID, &item.Name, &item.Notes, &item.Description, &item.Stage, &item.Category)

	return
}
