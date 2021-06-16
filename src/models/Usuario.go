package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido.
func (usuario *Usuario) Preparar(etapa string) error {

	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {

	if usuario.Nome == "" {
		return errors.New("o nome do usuário não pode ser em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o nick do usuário não pode ser em branco")
	}
	if usuario.Email == "" {
		return errors.New("o E-mail do usuário é obirgatório e não pode estar em branco")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o e-mail inserido é invalido")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("a senha do usuário não pode ser em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
