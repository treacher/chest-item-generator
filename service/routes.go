package service

import (
	"chest-item-generator/handlers"
	"chest-item-generator/logic"
	"net/http"

	"github.com/pressly/chi"
)

func Routes(config *Config) http.Handler {
	router := chi.NewRouter()
	itemRoller := logic.NewItemRoller(config.ItemMap)

	router.Get("/chest_items", handlers.NewGetChestItems(itemRoller).ServeHTTP)

	return router
}
