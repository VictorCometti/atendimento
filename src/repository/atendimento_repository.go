package repository

import (
	"database/sql"
	"log"

	"github.com/VictorCometti/atendimento.git/src/model"
)

func NewRepositoryAtendimentoRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repository Repository) GetAtendimento(cdAtendimento int) (atendimento model.Atendimento, erro error) {
	query := "SELECT descricao_atendimento, status_atendimento FROM atendimentos WHERE cd_atendimento = $1"

	rows, erro := repository.DB.Query(query, cdAtendimento)
	if erro != nil {
		log.Fatalf("Erro ao tentar executar a query. Erro: %v", erro)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Atendimento
		if erro = rows.Scan(
			&a.DescricaoAtendimento,
			&a.StatusAtendimento,
		); erro != nil {
			log.Fatalf("Erro ao tentar escanear as linhas. Erro: %v", erro)
		}
		atendimento = a
	}
	return

}
