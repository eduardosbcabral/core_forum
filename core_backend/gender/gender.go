package gender

type Gender struct {
	Id			int64	`json:"id"`
	Category 	string	`json:"category" validate:"required"`
}