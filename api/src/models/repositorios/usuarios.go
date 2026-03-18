package repositorios

import (
	"apiDevbook/src/models"
	"database/sql"
	"log"
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
		log.Fatal(err)
	}
	defer statement.Close()

	result, err := statement.Exec(usuario.Nome, usuario.Nick,
		usuario.Email, usuario.Password)
	if err != nil {
		return 0, err
	}

}
