package repository

import "go-crud/types"

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{ // 초기화
			{
				Name: "abc",
				Age:  3,
			},
			{
				Name: "test",
				Age:  10,
			},
		},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Update(beforeUser *types.User, updatedUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
