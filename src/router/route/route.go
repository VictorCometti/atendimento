package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI    string
	Method string
	Func   func(w http.ResponseWriter, r *http.Request)
}

func ConfigurationRouter(r *mux.Router) *mux.Router {
	rotas := rotaAtendimento

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Func).Methods(rota.Method)
	}

	return r
}

func ConfigurationPacienteRoute(r *mux.Router) *mux.Router {
	rotas := rotaPaciente

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Func).Methods(rota.Method)
	}
	return r
}
