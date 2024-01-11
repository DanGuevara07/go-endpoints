package users

import (
	"encoding/json"
	"fmt"
	"go-endpoints/internal/infrastructure/db"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users = new(db.Users)
	db.DB().Limit(10).Find(&users)

	return c.JSON(http.StatusOK, fromDBUsersToDTO(users))
}

func GetUser(c echo.Context) error {
	userId := c.Param("id")
	var user = db.User{ID: uuid.Must(uuid.Parse(userId))}
	db.DB().First(&user)

	return c.JSON(http.StatusOK, fromDBUserToDTO(&user))
}

func UpdateUser(c echo.Context) error {
	userId := c.Param("id")
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading request body: %s", err))
	}

	var userDTO UserDTO
	err = json.Unmarshal(body, &userDTO)

	var user = db.User{ID: uuid.Must(uuid.Parse(userId))}
	db.DB().First(&user)

	user.Name = userDTO.Name

	db.DB().Save(&user)

	return c.JSON(http.StatusOK, fromDBUserToDTO(&user))
}
func CreateUser(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading request body: %s", err))
	}

	var userDTO UserDTO
	err = json.Unmarshal(body, &userDTO)
	var userDB = new(db.User)
	userDB.Name = userDTO.Name
	userDB.ID = uuid.New()
	db.DB().Create(userDB)

	userDTO.CreatedAt = userDB.CreatedAt
	userDTO.UpdatedAt = userDB.UpdatedAt
	userDTO.ID = userDB.ID

	return c.JSON(http.StatusOK, userDTO)
}

func DeleteUser(c echo.Context) error {
	userId := c.Param("id")
	var user = db.User{ID: uuid.Must(uuid.Parse(userId))}

	db.DB().Delete(&user)

	return c.NoContent(204)
}
