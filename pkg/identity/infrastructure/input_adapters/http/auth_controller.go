package http

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/core/errors"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/application"
	dtos_http "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/input_adapters/http/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerType struct {
	authService application.IAuthService
}

func AuthController(authService application.IAuthService) *AuthControllerType {
	return &AuthControllerType{authService}
}

// Login
// @Summary		Iniciar sesión
// @Description	Autentica a un usuario y devuelve un token JWT
// @Tags		Identity
// @Accept		json
// @Produce		json
// @Param		credentials	body		dtos_http.LoginDto	true	"LoginDto"
// @Success		200			{object}	map[string]string
// @Failure		401
// @Router			/api/identity/login [post]
func (h *AuthControllerType) Login(c *gin.Context) {
	var creds dtos_http.LoginDto
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := h.authService.Login(creds)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register
// @Summary		Register User
// @Description	Registra un nuevo usuario
// @Tags		Identity
// @Accept		json
// @Produce		json
// @Param		credentials	body		dtos_http.RegisterDto	true	"RegisterDto"
// @Success		200			{object}	map[string]string
// @Router			/api/identity/register [post]
func (h *AuthControllerType) Register(c *gin.Context) {
	var user dtos_http.RegisterDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Error: err.Error()})
		return
	}

	err := h.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}
