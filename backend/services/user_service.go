package services

import (
	"context"

	"github.com/WiInfFelix/notr/ent"
	"github.com/WiInfFelix/notr/models"
)

type UserService struct {
	Client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		Client: client,
	}
}

func (s *UserService) GetAllUsers(page, itemsPerPage int) ([]*ent.User, error) {
	return s.Client.User.Query().Offset((page - 1) * itemsPerPage).Limit(itemsPerPage).All(context.Background())
}

func (s *UserService) CreateUser(userInput models.NewUserInput) (*ent.User, error) {
	return s.Client.User.Create().SetUsername(userInput.Username).SetPassword(userInput.Password).SetEmail(userInput.Email).Save(context.Background())
}
