package domain

import (
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type IUsuario interface {
	Save(usuario *entities.Usuario) error
	Update(usuario *entities.Usuario) error
	Delete(id int32) error
	GetById(id int32) (*entities.Usuario, error)
	GetAll() ([]entities.Usuario, error)
	GetByEmail(email string) (*entities.Usuario, error)
	GetByUsername(username string) (*entities.Usuario, error)
}
