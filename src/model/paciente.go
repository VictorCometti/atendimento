package model

import "time"

type Paciente struct {
	ID                int       `json:"-"`
	CodigoPaciente    int       `json:"cd_paciente,omitempty"`
	NomePaciente      string    `json:"nome_paciente,omitempty"`
	SobrenomePaciente string    `json:"sobrenome_paciente,omitempty"`
	CpfPaciente       string    `json:"cpf_paciente,omitempty"`
	RgPaciente        string    `json:"rg_paciente,omitempty"`
	DataNascimento    time.Time `json:"data_nascimento,omitempty"`
	NomeMae           string    `json:"nome_mae"`
	NomePai           string    `json:"nome_pai"`
}
