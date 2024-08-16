package repository

import (
	"sync"
)

// 3티어 아키텍처
// -> 세션 정의

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	// DB 연결
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})
	return repositoryInstance
}
