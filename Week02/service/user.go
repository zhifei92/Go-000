package service

import(
	"github.com/pkg/errors"
	"dao"
)

type User struct{

}

func NewUser() *Service {
	return &User{}
}

func (u *User) Get(id int) (dao.User, error) {
	u, err := dao.GetUserById(id)
	return u, err
}
