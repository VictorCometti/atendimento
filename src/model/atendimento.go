package model

type Atendimento struct {
	ID                   int      `json:"id,omitempty"`
	CodigoAtendimento    int      `json:"cd_atendimento,omitempty"`
	DescricaoAtendimento string   `json:"descricao_atendimento,omitempty"`
	Paciente             Paciente `json:"-"`
	StatusAtendimento    string   `json:"status_atendimento,omitempty"`
}
