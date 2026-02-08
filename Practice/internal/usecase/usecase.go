package usecase

import "fmt"

type UserUsecase struct {
}

func NewUserUsecase() UserUsecase {
	return UserUsecase{}
}

func (u UserUsecase) CreateUser(name string) string {
	fmt.Println(name)
	result := fmt.Sprintf("Hello %s", name)

	return result
}
