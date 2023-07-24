package user

type UserDto struct {
	Id       int64  `json:"id"`
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Username string `validate:"required,min=1,max=200" json:"username"`
	Password string `validate:"required,min=1,max=200" json:"password"`
}

func NewUserDto(id int64, name string, username string, password string) *UserDto {
	return &UserDto{
		Id:       id,
		Name:     name,
		Username: username,
		Password: password,
	}
}

func NewUserFromModel(model *User) *UserDto {
	return &UserDto{
		Id:       model.Id,
		Name:     model.Name,
		Username: model.Username,
		Password: model.Password,
	}
}
