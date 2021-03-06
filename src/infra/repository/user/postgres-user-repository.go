package user

import (
	"gtodo/src"
	"gtodo/src/entity/user"
	"gtodo/src/infra/client"
	gError "gtodo/src/infra/error"
)

type UserRepository struct {
	mapper       *UserMapper
	db           *client.DB
	errorFactory src.IErrorFactory
}

func (this *UserRepository) Create(user *user.User) (*user.User, error) {
	result := this.db.PConn.Create(&user)

	if result.Error != nil {
		return nil, this.errorFactory.BadRequestError("", result.Error)
	}

	return user, nil
}

func (this *UserRepository) Find(options interface{}) (*[]user.User, error) {
	users := new([]user.User)
	result := this.db.PConn.Find(&users, options)

	if result.Error != nil {
		return nil, this.errorFactory.InternalServerError(src.FIND_USER_ERROR, result.Error)
	}

	return users, nil
}

func (this *UserRepository) FindOne(options interface{}) (*user.User, error) {
	user := new(user.User)
	result := this.db.PConn.First(&user, options)

	if result.Error != nil {
		return nil, this.errorFactory.InternalServerError(src.FIND_ONE_USER_ERROR, result.Error)
	}

	return user, nil
}

func (this *UserRepository) UpdateById(id string, user *user.User) (*user.User, error) {
	return nil, nil
}

func (this *UserRepository) DeleteById(id string) (bool, error) {
	return true, nil
}

func NewUserRepository() user.IUserRepository {
	return &UserRepository{
		&UserMapper{},
		client.NewDB(),
		gError.NewErrorFactory(),
	}
}
