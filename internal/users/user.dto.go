package users

import (
	"encoding/json"
	"go-endpoints/internal/infrastructure/db"
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func fromDBUserToDTO(user *db.User) *UserDTO {
	userJSON, _ := json.Marshal(user)

	var userDTO = new(UserDTO)

	json.Unmarshal(userJSON, &userDTO)

	return userDTO
}

func fromDBUsersToDTO(users *db.Users) *[]UserDTO {
	usersJSON, _ := json.Marshal(users)

	var usersDTO = new([]UserDTO)

	json.Unmarshal(usersJSON, &usersDTO)

	return usersDTO
}
