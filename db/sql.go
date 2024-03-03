package db

import (
	"database/sql"
	"encoding/json"
)

func CreateClient() (db *sql.DB, err error) {
	return sql.Open("sqlite3", "./example.db")
}

func GetAllItems() (rows *sql.Rows, err error) {
	db, err := CreateClient()
	if err != nil {
		return
	}

	defer db.Close()

	rows, err = db.Query("SELECT * FROM items")
	if err != nil {
		return
	}

	// fmt.Println("rows", rows)

	defer rows.Close()
	return
}

func GetAll(table EntityLabel) (rows *sql.Rows, err error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err = db.Query("SELECT * FROM ?", table)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return
}

func GetByID(table EntityLabel, id int) (row *sql.Row, err error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	row = db.QueryRow("SELECT * FROM ? WHERE id = ?", table, id)
	return
}

func Put[T Entity](table EntityLabel, record Entity) (sql.Result, error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := `UPDATE ? SET name = ?, notes = ?, description = ?, stage = ?, category = ? WHERE id = ?`

	result, err := db.Exec(query, table, record.Name, record.Notes, record.Description, record.Stage, record.Category, record.ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func PostItem(record Item) (sql.Result, error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := `INSERT INTO items (name, notes, description, stage, category)
	VALUES (?, ?, ?, ?, ?)`

	result, err := db.Exec(query, record.Name, record.Notes, record.Description, record.Stage, record.Category)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func PostBox(record Box) (sql.Result, error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := `INSERT INTO boxes (name, notes, description, stage, category)
	VALUES (?, ?, ?, ?, ?)`

	result, err := db.Exec(query, record.Name, record.Notes, record.Description, record.Stage, record.Category)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Delete(table EntityLabel, id int) (sql.Result, error) {
	db, err := CreateClient()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := `DELETE FROM ? WHERE id = ?`

	return db.Exec(query, table, id)
}

func ParseItem(item *Item, scan func(dest ...any) error) (err error) {
	return scan(&item.ID, &item.Name, &item.Notes, &item.Description, &item.Stage, &item.Category)
}

func ParseBox(box *Box, scan func(dest ...any) error) error {
	return scan(&box.ID, &box.Name, &box.Notes, &box.Description, &box.Stage, &box.Category)
}

func ParseEntityFromBytes(b []byte) (entity Entity, err error) {
	err = json.Unmarshal(b, &entity)
	return
}

func ParseItemFromBytes(b []byte) (item Item, err error) {
	err = json.Unmarshal(b, &item)
	return
}

func ParseBoxFromBytes(b []byte) (box Box, err error) {
	err = json.Unmarshal(b, &box)
	return
}
