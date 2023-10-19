package users

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName" validate:"min=2,max=64,required"`
	LastName  string `json:"lastName" validate:"min=2,max=64,required"`
	Email     string `json:"email" validate:"required,email"`
}
