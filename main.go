package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
)

var ItemMap = make(map[string][]Item)

const (
	COMMON     = "common"
	RARE       = "rare"
	SUPER_RARE = "super"
	ULTRA_RARE = "ultra"
	ANCIENT    = "ancient"
)

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

func getRandomNumber(max int64) int64 {
	number, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		println(err)
	}

	return number.Int64()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getRareItem(rarity string) Item {
	items := ItemMap[rarity]
	return items[getRandomNumber(int64(len(items)))]
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var chestItems []ItemGroup

		chestItems = append(chestItems, rollForCommonItems()...)
		chestItems = append(chestItems, rollForRareItems()...)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chestItems)
	}
}

func main() {
	file, e := ioutil.ReadFile("./items.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &ItemMap)

	http.HandleFunc("/chest_items", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", getenv("PORT", "8080")), nil)
}

func rollForCommonItems() []ItemGroup {
	var commonChestItems []ItemGroup

	commonItems := ItemMap[COMMON]

	for _, item := range commonItems {
		quantity := getRandomNumber(int64(item.MaxQty))
		if quantity == 0 {
			continue
		}
		itemGroup := ItemGroup{Item: item, Quantity: quantity}
		commonChestItems = append(commonChestItems, itemGroup)
	}

	return commonChestItems
}

func rollForRareItems() []ItemGroup {
	diceRoll := getRandomNumber(100)

	var rareItem Item

	switch {
	case diceRoll > 65 && diceRoll <= 85:
		rareItem = getRareItem(RARE)
	case diceRoll > 85 && diceRoll <= 95:
		rareItem = getRareItem(SUPER_RARE)
	case diceRoll > 95 && diceRoll <= 99:
		rareItem = getRareItem(ULTRA_RARE)
	case diceRoll == 100:
		rareItem = getRareItem(ANCIENT)
	}

	if rareItem != (Item{}) {
		return []ItemGroup{{Item: rareItem, Quantity: 1}}
	}

	return []ItemGroup{}
}
