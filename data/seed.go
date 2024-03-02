package data

func getSeedData() (items []Item, boxes []Box, boxitems []BoxItem) {
	items = []Item{
		{
			ID: 1,
			Name: "Toothbrush",
			Stage: Essentials,
			Category: Bathroom,
		},
		{
			ID: 2,
			Name: "Toothpaste",
			Stage: Essentials,
			Category: Bathroom,
		},
		{
			ID: 3,
			Name: "TV",
			Stage: StageTwo,
			Category: Bedroom,
		},
		{
			ID: 4,
			Name: "Micro USB Bundle",
			Stage: StageOne,
			Category: Office,
		},
	}

	plasticTubDescription := "Plastic tub with blue lid"

	boxes = []Box{
		{
			ID: 1,
			Name: "Cable Box",
			Description: &plasticTubDescription,
			Stage: StageOne,
		},
	}

	boxitems = []BoxItem{
		{
			ID: 1,
			BoxID: 1,
			ItemID: 4,
		},
	}

	return
}
