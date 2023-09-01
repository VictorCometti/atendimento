package model

import "time"

type Exame struct {
	ID                int         `json:"id,omitempty"`
	CodigoExame       int         `json:"codigo_exame,omitempty"`
	DescricaoExame    string      `json:"descricao_exame,omitempty"`
	DataExame         time.Time   `json:"data_exame,omitempty"`
	MedicoSolicitante Medico      `json:"cd_medico,omitempty"`
	CodigoAtendimento Atendimento `json:"cd_atendimento,omitempty"`
}
