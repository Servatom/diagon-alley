package repository_auth

import domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"

func (userRepo UserRepository) toDomainUser() *domain_auth.UserWithID {
	return &domain_auth.UserWithID{
		ID:       userRepo.ID,
		User:     domain_auth.User{
			Email:    userRepo.Email,
			Password: userRepo.Password,
			Gender: userRepo.Gender,
			Age: userRepo.Age,
		},
	}
}
func NewUserRepository(
	user *domain_auth.User,
) *UserRepository {
	return &UserRepository{
		Email:        user.Email,
		Password:     user.Password,
		Gender: user.Gender,
		Age: user.Age,
	}
}
