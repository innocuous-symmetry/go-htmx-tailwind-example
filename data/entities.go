package data

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

// entities
type Item struct {
	ID    			int
	Name  			string
	Notes 			*string
	Description		*string
	Stage 			PackingStage
	Category		Category
}

type Box struct {
	ID				int
	Name			string
	Notes			*string
	Description		*string
	Stage			PackingStage
	Category		Category
}

// joins
type BoxItem struct {
	ID				int
	BoxID			int
	ItemID			int
}
