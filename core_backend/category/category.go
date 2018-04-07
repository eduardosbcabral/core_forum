package category

type Category struct {
	Id			int64	`json:"id"`
	Category 	string	`json:"category" validate:"required"`
}