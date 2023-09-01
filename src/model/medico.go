package model

type Medico struct {
	ID             int    `json:"id,omitempty"`
	NomeMedico     string `json:"nome_medico"`
	TelefoneMedico string `json:"telefone_medico,omitempty"`
	Crm            string `json:"crm,omitempty"`
	CodigoMedico   int    `json:"cd_medico,omitempty"`
}
