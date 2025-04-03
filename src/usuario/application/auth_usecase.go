package application

import (
	"errors"

	"github.com/vicpoo/NetflixAPIgo/src/core/auth"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type AuthUseCase struct {
	repo domain.IUsuario
}

func NewAuthUseCase(repo domain.IUsuario) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) Login(email, password string) (string, *entities.Usuario, error) {
	if email == "" || password == "" {
		return "", nil, errors.New("email y contraseña son requeridos")
	}

	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (uc *AuthUseCase) LoginWithUsername(username, password string) (string, *entities.Usuario, error) {
	if username == "" || password == "" {
		return "", nil, errors.New("nombre de usuario y contraseña son requeridos")
	}

	user, err := uc.repo.GetByUsername(username)
	if err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
