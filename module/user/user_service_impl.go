package user

import (
	"denitiawan/research-swagger-gomod-gin/common/dto"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepo UserRepo
	Validate *validator.Validate
	Class    string
}

func NewUserServiceImpl(repo UserRepo, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: repo,
		Validate: validate,
		Class:    "UserServiceImpl",
	}
}

func (t *UserServiceImpl) Create(requestDto UserDto) *dto.ImplResponse {
	function := "Create()"

	err := t.Validate.Struct(requestDto)
	if err != nil {
		return dto.NewImplResponse(t.Class, function, "validate dto failed", err, nil)
	}

	// model
	model := NewUser(0, requestDto.Name, requestDto.Username, requestDto.Password)

	// repo
	res := t.UserRepo.Save(model)
	if res.Error != nil {
		return dto.NewImplResponse(res.Class, res.Function, res.Message, res.Error, nil)
	}

	// ok
	return dto.NewImplResponse(t.Class, function, res.Message, nil, res.Data)

}

func (t *UserServiceImpl) Delete(id int64) *dto.ImplResponse {
	function := "FindById()"

	// find
	res := t.UserRepo.FindById(id)
	if res.Error != nil {
		return dto.NewImplResponse(res.Class, res.Function, res.Message, res.Error, nil)
	}

	// delete
	res = t.UserRepo.Delete(id)
	if res.Error != nil {
		return dto.NewImplResponse(res.Class, res.Function, res.Message, res.Error, nil)
	}

	// ok
	return dto.NewImplResponse(t.Class, function, res.Message, nil, res.Data)

}

func (t *UserServiceImpl) FindAll() *dto.ImplResponse {
	function := "FindAll()"

	// repo
	res := t.UserRepo.FindAll()
	if res.Error != nil {
		return dto.NewImplResponse(res.Class, res.Function, res.Message, res.Error, nil)
	}

	// ok
	return dto.NewImplResponse(t.Class, function, res.Message, nil, res.Data)
}

func (t *UserServiceImpl) FindById(id int64) *dto.ImplResponse {
	function := "FindById()"

	// repo
	res := t.UserRepo.FindById(id)
	if res.Error != nil {
		return dto.NewImplResponse(res.Class, res.Function, res.Message, res.Error, nil)
	}

	// ok
	return dto.NewImplResponse(t.Class, function, res.Message, nil, res.Data)

}

func (t *UserServiceImpl) Update(id int64, requestDto UserDto) *dto.ImplResponse {
	function := "Update()"

	// find
	res := t.UserRepo.FindById(id)
	if res.Error != nil {
		return dto.NewImplResponse(res.Class, res.Function, res.Message, res.Error, nil)
	}

	// set new value
	data, _ := res.Data.(User)
	data.Name = requestDto.Name
	data.Username = requestDto.Username
	data.Password = requestDto.Password

	// update
	res = t.UserRepo.Update(data)
	if res.Error != nil {
		return dto.NewImplResponse(res.Class, res.Function, res.Message, res.Error, nil)
	}

	// ok
	return dto.NewImplResponse(t.Class, function, res.Message, nil, data)
}
