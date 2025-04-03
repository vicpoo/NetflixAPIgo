package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/application"
)

type DeleteUsuarioController struct {
	deleteUseCase *application.DeleteUsuarioUseCase
}

func NewDeleteUsuarioController(deleteUseCase *application.DeleteUsuarioUseCase) *DeleteUsuarioController {
	return &DeleteUsuarioController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteUsuarioController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el usuario",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Usuario eliminado exitosamente",
	})
}
