package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositorio de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes ( titulo, conteudo, autor_id ) VALUES ( ?, ?, ? )",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// BuscarPorID traz uma publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(ID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		SELECT
			P.id, 
			P.titulo, 
			P.conteudo, 
			P.autor_id, 
			U.nick,
			P.curtidas,
			P.criadaEm
		FROM
			publicacoes P INNER JOIN
				usuarios U ON U.id = P.autor_id
		WHERE
			P.id = ?`,
		ID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuarios seguidos e tambem do proprio usuario que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT
			DISTINCT
			P.id, 
			P.titulo, 
			P.conteudo, 
			P.autor_id, 
			U.nick,
			P.curtidas,
			P.criadaEm
		FROM
			publicacoes P
				INNER JOIN usuarios U ON U.id = P.autor_id
				INNER JOIN seguidores S ON P.autor_id = S.usuario_id
		WHERE
			P.autor_id = ? OR
			S.seguidor_id = ?
		ORDER BY
			P.criadaEm DESC`,
		usuarioID,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera as informações de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(ID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM publicacoes WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUsuario traz todas as publicações de um usuario especifico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT
			P.id, 
			P.titulo, 
			P.conteudo, 
			P.autor_id, 
			U.nick,
			P.curtidas,
			P.criadaEm
		FROM
			publicacoes P
				INNER JOIN usuarios U ON U.id = P.autor_id
		WHERE
			P.autor_id = ?
		ORDER BY
			P.criadaEm DESC`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida na publicação
func (repositorio Publicacoes) Descurtir(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET curtidas = GREATEST(curtidas - 1, 0) WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
