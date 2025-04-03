package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/application"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type CreateUsuarioController struct {
	createUseCase *application.CreateUsuarioUseCase
}

func NewCreateUsuarioController(createUseCase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateUsuarioController) Run(c *gin.Context) {
	var usuarioRequest struct {
		Name     string `json:"name" binding:"required"`
		Lastname string `json:"lastname" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&usuarioRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	usuario := entities.NewUsuario(
		usuarioRequest.Name,
		usuarioRequest.Lastname,
		usuarioRequest.Username,
		usuarioRequest.Password,
		usuarioRequest.Email,
	)

	createdUsuario, err := ctrl.createUseCase.Run(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUsuario)
}
