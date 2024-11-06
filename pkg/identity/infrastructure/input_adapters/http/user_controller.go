package http

import (
	errors "github.com/yasniel1408/hexa-ddd-golang-gin/core/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/application"
)

type UserControllerType struct {
	userService application.IUserService
}

func UserController(userService application.IUserService) *UserControllerType {
	return &UserControllerType{userService}
}

// GetUser
// @Summary		Obtiene un usuario por su ID
// @Description	Devuelve los datos de un usuario específico basado en el ID proporcionado
// @Tags		Identity
// @Param		id	path		int	true	"ID del Usuario"
// @Success		200	{object}	dtos_http.GetUserResponseDto
// @Failure		400	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/api/identity/users/{id} [get]
func (h *UserControllerType) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Error: "invalid user id"})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, errors.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
