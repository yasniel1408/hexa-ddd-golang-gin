package in

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/application"
)

type UserHandler struct {
    userService application.UserService
}

func NewUserHandler(userService application.UserService) *UserHandler {
    return &UserHandler{userService}
}

// GetUser obtiene un usuario por su ID
func (h *UserHandler) GetUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }

    user, err := h.userService.GetUserByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}