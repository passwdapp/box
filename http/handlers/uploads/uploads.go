package uploads

import "github.com/passwdapp/box/config"

// Handler contains all the handlers in the uploads package
type Handler struct {
	Config *config.Config
}

// Init initializes the Handler struct
func (h *Handler) Init(cfg *config.Config) {
	h.Config = cfg
}
