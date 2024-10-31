package in

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/auth/application"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/auth/domain/valueobjects"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
)

// AuthHandler maneja las solicitudes de autenticación
type AuthHandler struct {
    authService application.AuthService
}

func NewAuthHandler(authService application.AuthService) *AuthHandler {
    return &AuthHandler{authService}
}

// Login realiza la autenticación de un usuario
// @Summary Iniciar sesión
// @Description Autentica a un usuario y devuelve un token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body valueobjects.Credentials true "Credenciales"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    var creds valueobjects.Credentials
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := h.authService.Login(creds)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register registra un nuevo usuario
func (h *AuthHandler) Register(c *gin.Context) {
    var user entities.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.authService.Register(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}