package router

import (
	"github.com/VictorCometti/atendimento.git/src/router/route"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	route.ConfigurationRouter(r)
	route.ConfigurationPacienteRoute(r)
	return r
}
