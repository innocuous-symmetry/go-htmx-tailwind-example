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

// entities
type Item struct {
	ID    			int				`json:"id"`
	Name  			string			`json:"name"`
	Notes 			*string			`json:"notes"`
	Description		*string			`json:"description"`
	Stage 			PackingStage	`json:"stage"`
	Category		Category		`json:"category"`
}

type Box struct {
	ID				int				`json:"id"`
	Name			string			`json:"name"`
	Notes			*string			`json:"notes"`
	Description		*string			`json:"description"`
	Stage			PackingStage	`json:"stage"`
	Category		Category		`json:"category"`
}

// joining tables and derivative data types
type BoxItem struct {
	ID				int
	BoxID			int
	ItemID			int
}

type BoxItemWithItemInfo struct {
	ID				int
	Name			string
	Stage 			PackingStage
	Category		Category
	Description		*string
	Notes			*string
}
