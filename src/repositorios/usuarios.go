package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// usuarios representa um repositorio de usuários.
type usuarios struct {
	db *sql.DB
}

// NovoRepositoriosDeUsuarios cria um novo repositório de usuário
func NovoRepositoriosDeUsuarios(db *sql.DB) *usuarios {

	return &usuarios{db}
}

func (repo usuarios) Criar(usuario models.Usuario) (uint64, error) {

	statement, erro := repo.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha)VALUES(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar busca todos os usuários que atendem um filtro de nome ou nick
func (repositorio usuarios) Buscar(filtro string) ([]models.Usuario, error) {

	nomeOuNick := fmt.Sprintf("%%%s%%", filtro)

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoem FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID busca um usuário pelo ID di banco
func (repositorio usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoem FROM usuarios WHERE id = ?",
		ID,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar atualiza os dados do usuário no banco
func (repositorio usuarios) Atualizar(ID uint64, usuario models.Usuario) error {

	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar deleta o usuário pelo ID do banco
func (repositorio usuarios) Deletar(ID uint64) error {

	statement, erro := repositorio.db.Prepare(
		"DELETE FROM usuarios WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu ID e a senha hash
func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error) {

	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}
