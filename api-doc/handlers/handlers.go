package handlers

import (
	"net/http"

	"github.com/Shaheer25/api-doc/models"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Data []models.User `json:"data"`
}

// GetUsers return list of all users from the Database
// @Summarry return list of all
// @Description return list of all users in the Database
// @Tags Users
// @Success 200 {object} userResponse
// @Router /users [get]

func GetUsers(c *gin.Context) {

	users := models.ListUser()

	c.JSON(http.StatusOK, userResponse{Data: users})
}
