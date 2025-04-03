package application

import (
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type GetAllUsuariosUseCase struct {
	repo domain.IUsuario
}

func NewGetAllUsuariosUseCase(repo domain.IUsuario) *GetAllUsuariosUseCase {
	return &GetAllUsuariosUseCase{repo: repo}
}

func (uc *GetAllUsuariosUseCase) Run() ([]entities.Usuario, error) {
	usuarios, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}
