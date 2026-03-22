package repositorios

import (
	"apiDevbook/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepoPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repo Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repo.db.Prepare(`
		insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publicacao.Titulo,
		publicacao.Conteudo,
		publicacao.AutorId)
	if err != nil {
		return 0, err
	}

	ultimoIdInserido, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIdInserido), nil
}

func (repo Publicacoes) BuscarId(publicacaoId uint64) (models.Publicacao, error) {
	query, err := repo.db.Query(`
		select p.*, u.nick from publicacoes p join usuarios u on 
		p.autor_id = u.id where p.id = ?
	`, publicacaoId)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer query.Close()
	var publicacao models.Publicacao

	if query.Next() {

		if err := query.Scan(&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.DataCriacao,
			&publicacao.AutorNick); err != nil {
			return models.Publicacao{}, err
		}
	}

	return publicacao, nil
}

func (repo Publicacoes) Buscar(userId uint64) ([]models.Publicacao, error) {
	query, err := repo.db.Query(`
		select distinct p.*, u.nick from
		publicacoes p inner join usuarios u on u.id = p.autor_id
		left join seguidores s on p.autor_id = s.usuario_id
		where  u.id = ? or s.seguidor_id = ?
		order by 1 desc
	`, userId, userId)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var publicacoes []models.Publicacao

	for query.Next() {
		var publicacao models.Publicacao
		if err := query.Scan(&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.DataCriacao,
			&publicacao.AutorNick); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}
