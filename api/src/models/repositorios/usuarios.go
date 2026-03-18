package repositorios

import (
	"apiDevbook/src/models"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// Criar um repositorio de usuarios
func NovoRepoUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repo Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(usuario.Nome, usuario.Nick,
		usuario.Email, usuario.Password)
	if err != nil {
		return 0, err
	}

	ultIdInserido, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultIdInserido), nil

}

func (repo Usuarios) Buscar(nickOrName string) ([]models.Usuario, error) {
	nickOrName = fmt.Sprintf("%%%s%%", nickOrName)

	query, err := repo.db.Query(
		"select id, nome, nick, email, data_inicio from usuarios where nome like ? or nick like ?",
		nickOrName, nickOrName)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var usuarios []models.Usuario

	for query.Next() {
		var usuario models.Usuario
		if err = query.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.DataInicio,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}
