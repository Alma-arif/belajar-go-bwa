package handler

import (
	"membuat-api-bwa/helper"
	"membuat-api-bwa/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorsMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, _ := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account - failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentoken")

	response := helper.APIResponse("Account has been reegistered", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)

}
