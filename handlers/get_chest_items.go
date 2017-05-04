package handlers

import (
	"encoding/json"
	"net/http"
	"personal/chest-item-generator/logic"
)

type GetChestItems struct {
	itemRoller *logic.ItemRoller
}

func NewGetChestItems(itemRoller *logic.ItemRoller) *GetChestItems {
	return &GetChestItems{
		itemRoller: itemRoller,
	}
}

func (gci *GetChestItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chestItems := gci.itemRoller.GetChestItems()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chestItems)
}
