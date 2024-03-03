package db

// enums
type PackingStage int
type Category int

const (
	Essentials PackingStage = iota
	StageOne
	StageTwo
	StageThree
)

const (
	Bedroom Category = iota
	Bathroom
	Kitchen
	Office
	LivingRoom
	Other
)

var CategoryMap = map[Category]string{
	Bedroom:    "Bedroom",
	Bathroom:   "Bathroom",
	Kitchen:    "Kitchen",
	Office:     "Office",
	LivingRoom: "Living Room",
	Other:      "Other",
}

var PackingStageMap = map[PackingStage]string{
	Essentials: "Essentials",
	StageOne:   "Stage One",
	StageTwo:   "Stage Two",
	StageThree: "Stage Three",
}

type EntityLabel string

const (
	ItemType    EntityLabel = "items"
	BoxType     EntityLabel = "boxes"
	BoxItemType EntityLabel = "box_items"
)

type Entity struct {
	ID          	int
	EntityLabel 	EntityLabel
	Name        	string
	Notes       	*string
	Description 	*string
	Stage       	PackingStage
	Category    	Category
}

type Item Entity
type Box Entity

// joins
type BoxItem struct {
	ID     int
	BoxID  int
	ItemID int
}
