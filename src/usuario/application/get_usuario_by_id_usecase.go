package application

import (
	"errors"

	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type GetUsuarioByIdUseCase struct {
	repo domain.IUsuario
}

func NewGetUsuarioByIdUseCase(repo domain.IUsuario) *GetUsuarioByIdUseCase {
	return &GetUsuarioByIdUseCase{repo: repo}
}

func (uc *GetUsuarioByIdUseCase) Run(id int32) (*entities.Usuario, error) {
	if id <= 0 {
		return nil, errors.New("ID de usuario inválido")
	}

	usuario, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}
