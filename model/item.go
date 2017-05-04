package model

import "fmt"

type Item struct {
	Identifier  int64  `json:"identifier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      string `json:"rarity"`
	MaxQty      int    `json:"maxQty"`
}

type ItemGroup struct {
	Item     Item  `json:"item"`
	Quantity int64 `json:"quantity"`
}

func (itemGroup ItemGroup) String() string {
	return fmt.Sprintf("Item Name: %s, Max Quantity: %d", itemGroup.Item.Name, itemGroup.Item.MaxQty)
}

func (item Item) String() string {
	return fmt.Sprintf("Item Name: %s, Max Quantity: %d", item.Name, item.MaxQty)
}
