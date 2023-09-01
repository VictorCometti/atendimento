package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VictorCometti/atendimento.git/src/database"
	"github.com/VictorCometti/atendimento.git/src/repository"
)

func GetPaciente(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cpf := params.Get("cpf")

	db, erro := database.GetConnection()

	if erro != nil {
		log.Fatalf("Erro ao tentar pegar a conexão: Erro: %v", erro)
	}

	defer db.Close()

	repository := repository.NewRepositoryPacienteRepository(db)
	paciente, erro := repository.SearchPacienteByCpf(cpf)
	if erro != nil {
		log.Fatalf("Erro ao tentar buscar paciente. Erro: %v", erro)
	}
	json, erro := json.Marshal(paciente)
	if erro != nil {
		log.Fatalf("Erro ao tentar converter paciente para JSON: Erro: %v", erro)
	}

	w.Write(json)

}
