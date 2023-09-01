package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/VictorCometti/atendimento.git/src/database"
	"github.com/VictorCometti/atendimento.git/src/repository"
)

func GetAtendimentoByCodigoIdentificao(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("codigo")
	cdAtendimento, erro := strconv.Atoi(param)
	if erro != nil {
		log.Fatalf("Erro ao tentar converter para inteiro. Erro: %v", erro)
	}

	db, erro := database.GetConnection()

	if erro != nil {
		log.Fatalf("Erro ao tentar abrir a conexão com banco de dados. Erro: %v", erro)
	}

	defer db.Close()

	repository := repository.NewRepositoryAtendimentoRepository(db)
	atendimento, erro := repository.GetAtendimento(cdAtendimento)
	if erro != nil {
		log.Fatalf("Erro ao tentar pegar o atendimento. Erro: %v", erro)
	}

	json, erro := json.Marshal(atendimento)
	if erro != nil {
		log.Fatalf("Erro ao tentar transformar para json. Erro: %v", erro)
	}

	w.Write(json)
}
