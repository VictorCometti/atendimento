package route

import (
	"net/http"

	"github.com/VictorCometti/atendimento.git/src/handlers"
)

var rotaAtendimento = []Route{
	{
		URI:    "/atendimento",
		Method: http.MethodGet,
		Func:   handlers.GetAtendimentoByCodigoIdentificao,
	},
	{
		URI:    "/atendimento",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
	},
	{
		URI:    "/atendimento",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
	},
	{
		URI:    "/atendimento",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
	},
}
