package dto

type CreateUserDTO struct {
	ID       string
	Name     string
	Password string
}

type UpdateUserDTO struct {
	ID       *string
	Name     *string
	Password *string
	IsActive *bool
}
