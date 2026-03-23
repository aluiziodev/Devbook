package repositorios

import (
	"apiDevbook/src/models"
	"database/sql"
	"errors"
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

func (repo Usuarios) BuscarId(id uint64) (models.Usuario, error) {
	query, err := repo.db.Query(
		"select id, nome, nick, email, data_inicio from usuarios where id=?",
		id)
	if err != nil {
		return models.Usuario{}, err
	}
	defer query.Close()

	var usuario models.Usuario
	if query.Next() {
		if err = query.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.DataInicio,
		); err != nil {
			return models.Usuario{}, err
		}

	}

	if usuario.Nome == "" {
		return models.Usuario{}, errors.New("Usuario nao existe na base de dados")
	}

	return usuario, nil
}

func (repo Usuarios) Atualizar(id uint64, usuario models.Usuario) error {
	statement, err := repo.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); err != nil {
		return err
	}

	return nil

}

func (repo Usuarios) Deletar(id uint64) error {
	statement, err := repo.db.Prepare(
		"delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil

}

func (repo Usuarios) BuscarEmail(email string) (models.Usuario, error) {
	query, err := repo.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer query.Close()

	var usuario models.Usuario

	if query.Next() {
		if err = query.Scan(&usuario.Id, &usuario.Password); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repo Usuarios) Seguir(idSeguir, idSeguidor uint64) error {
	statement, err := repo.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(idSeguir, idSeguidor); err != nil {
		return err
	}

	return nil
}

func (repo Usuarios) DeixarSeguir(idSeguir, idSeguidor uint64) error {
	statement, err := repo.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(idSeguir, idSeguidor); err != nil {
		return err
	}

	return nil
}

func (repo Usuarios) BuscarSeguidores(id uint64) ([]models.Usuario, error) {
	query, err := repo.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.data_inicio 
		from usuarios as u join seguidores as s on u.id = s.seguidor_id
		where s.usuario_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var seguidores []models.Usuario

	for query.Next() {
		var seguidor models.Usuario
		if err = query.Scan(&seguidor.Id, &seguidor.Nome, &seguidor.Nick, &seguidor.Email, &seguidor.DataInicio); err != nil {
			return nil, err
		}
		seguidores = append(seguidores, seguidor)
	}
	return seguidores, nil
}

func (repo Usuarios) BuscarSeguindo(id uint64) ([]models.Usuario, error) {
	query, err := repo.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.data_inicio 
		from usuarios as u join seguidores as s on u.id = s.usuario_id
		where s.seguidor_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var seguidores []models.Usuario

	for query.Next() {
		var seguidor models.Usuario
		if err = query.Scan(&seguidor.Id, &seguidor.Nome, &seguidor.Nick, &seguidor.Email, &seguidor.DataInicio); err != nil {
			return nil, err
		}
		seguidores = append(seguidores, seguidor)
	}
	return seguidores, nil

}

func (repo Usuarios) BuscarSenha(id uint64) (string, error) {
	query, err := repo.db.Query(`select senha from usuarios where id = ?`, id)
	if err != nil {
		return "", err
	}
	defer query.Close()

	var usuario models.Usuario

	if query.Next() {
		if err := query.Scan(&usuario.Password); err != nil {
			return "", err

		}
	}

	return usuario.Password, nil
}

func (repo Usuarios) AtualizarSenha(id uint64, senha string) error {
	statement, err := repo.db.Prepare(`update usuarios set senha = ? where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(senha, id); err != nil {
		return err
	}

	return nil

}
