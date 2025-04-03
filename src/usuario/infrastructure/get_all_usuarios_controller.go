package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/application"
)

type GetAllUsuariosController struct {
	getAllUseCase *application.GetAllUsuariosUseCase
}

func NewGetAllUsuariosController(getAllUseCase *application.GetAllUsuariosUseCase) *GetAllUsuariosController {
	return &GetAllUsuariosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllUsuariosController) Run(c *gin.Context) {
	usuarios, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los usuarios",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}
