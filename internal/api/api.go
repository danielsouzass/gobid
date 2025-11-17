package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/danielsouzass/gobid/internal/services"
	"github.com/go-chi/chi/v5"
)

type API struct {
	Router      *chi.Mux
	UserService services.UserService
	Sessions    *scs.SessionManager
}
