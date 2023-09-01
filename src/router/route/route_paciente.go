package route

import (
	"net/http"

	"github.com/VictorCometti/atendimento.git/src/handlers"
)

var rotaPaciente = []Route{
	{
		URI:    "/paciente",
		Method: http.MethodGet,
		Func:   handlers.GetPaciente,
	},
}
