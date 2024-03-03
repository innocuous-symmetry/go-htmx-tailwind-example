package db

import (
	"database/sql"
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

func PostItem(item Item) (int64, error) {
	db, err := CreateClient()
	if err != nil {
		return -1, err
	}

	defer db.Close()

	query := `INSERT INTO items (name, notes, description, stage, category)
	VALUES (?, ?, ?, ?, ?)`

	result, err := db.Exec(query, item.Name, item.Notes, item.Description, item.Stage, item.Category)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}



func PostBox(box Box) (int64, error) {
	db, err := CreateClient()
	if err != nil {
		return -1, err
	}

	defer db.Close()

	query := `INSERT INTO boxes (name, notes, description, stage, category) VALUES (?, ?, ?, ?, ?)`

	result, err := db.Exec(query, box.Name, box.Notes, box.Description, box.Stage, box.Category)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func GetAllBoxes() (result []Box, err error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM boxes")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		box := Box{}

		err = rows.Scan(&box.ID, &box.Name, &box.Notes, &box.Description, &box.Stage, &box.Category)

		if err != nil {
			return nil, err
		}

		result = append(result, box)
	}

	return
}

// func PostBoxItem(itemid int, boxid int) (int64, error) {
// 	db, err := CreateClient()
// 	if err != nil {
// 		return -1, err
// 	}

// 	defer db.Close()
// 	// query :=
// }
