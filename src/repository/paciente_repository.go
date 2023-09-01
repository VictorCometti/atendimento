package repository

import (
	"database/sql"
	"log"

	"github.com/VictorCometti/atendimento.git/src/model"
)

func NewRepositoryPacienteRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repository Repository) SearchPacienteByCpf(cpf string) (paciente model.Paciente, erro error) {
	sql := "SELECT cd_paciente, nome_paciente, sobrenome_paciente, cpf_paciente, rg_paciente, data_nascimento, nome_mae, nome_pai FROM pacientes WHERE cpf_paciente = $1"

	rows, erro := repository.DB.Query(sql, cpf)

	if erro != nil {
		log.Fatalf("Erro ao tentar executar a query")
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Paciente

		if erro = rows.Scan(
			&p.CodigoPaciente,
			&p.NomePaciente,
			&p.SobrenomePaciente,
			&p.CpfPaciente,
			&p.RgPaciente,
			&p.DataNascimento,
			&p.NomeMae,
			&p.NomePai,
		); erro != nil {
			log.Fatalf("Erro ao tentar escanear a linha: Erro %v", erro)
		}
		paciente = p
	}
	return
}
