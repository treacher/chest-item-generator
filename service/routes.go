package service

import (
	"net/http"
	"personal/chest-item-generator/handlers"
	"personal/chest-item-generator/logic"

	"github.com/pressly/chi"
)

func Routes(config *Config) http.Handler {
	router := chi.NewRouter()
	itemRoller := logic.NewItemRoller(config.ItemMap)

	router.Get("/chest_items", handlers.NewGetChestItems(itemRoller).ServeHTTP)

	return router
}
