package db

import (
	"database/sql"
	"os"
)

func GetSeedData() (items []Item, boxes []Box, boxitems []BoxItem) {
	items = []Item{
		{
			ID:       1,
			Name:     "Toothbrush",
			Stage:    Essentials,
			Category: Bathroom,
		},
		{
			ID:       2,
			Name:     "Toothpaste",
			Stage:    Essentials,
			Category: Bathroom,
		},
		{
			ID:       3,
			Name:     "TV",
			Stage:    StageTwo,
			Category: Bedroom,
		},
		{
			ID:       4,
			Name:     "Micro USB Bundle",
			Stage:    StageOne,
			Category: Office,
		},
	}

	plasticTubDescription := "Plastic tub with blue lid"

	boxes = []Box{
		{
			ID:          1,
			Name:        "Cable Box",
			Description: &plasticTubDescription,
			Stage:       StageOne,
		},
	}

	boxitems = []BoxItem{
		{
			ID:     1,
			BoxID:  1,
			ItemID: 4,
		},
	}

	return
}

func CreateTables(client *sql.DB) (int64, error) {
	script, err := os.ReadFile("/home/mikayla/go/go-htmx-tailwind-example/db/seed.sql")
	if err != nil {
		panic(err)
	}

	result, err := client.Exec(string(script))
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}

func SeedDB() (int64, error) {
	client, err := CreateClient()
	if err != nil {
		return -1, err
	}

	defer client.Close()

	CreateTables(client)

	items, boxes, _ := GetSeedData()
	insertCount := 0

	// loop through initial items and run post request for each
	for i := range(items) {
		_, err := PostItem(items[i])
		if err != nil {
			return -1, err
		}
		insertCount++
	}

	for i := range(boxes) {
		_, err := PostBox(boxes[i])
		if err != nil {
			return -1, err
		}
		insertCount++
	}

	return int64(insertCount), nil
}
