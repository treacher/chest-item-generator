package logic

import (
	"chest-item-generator/model"
	"crypto/rand"
	"math/big"
)

const (
	COMMON     = "common"
	RARE       = "rare"
	SUPER_RARE = "super"
	ULTRA_RARE = "ultra"
	ANCIENT    = "ancient"
)

type ItemRoller struct {
	ItemMap map[string][]model.Item
}

func NewItemRoller(itemMap map[string][]model.Item) *ItemRoller {
	return &ItemRoller{
		ItemMap: itemMap,
	}
}

func (ir *ItemRoller) GetChestItems() []model.ItemGroup {
	var chestItems []model.ItemGroup

	chestItems = append(chestItems, ir.rollForCommonItems()...)
	chestItems = append(chestItems, ir.rollForRareItems()...)

	return chestItems
}

func (ir *ItemRoller) rollForCommonItems() []model.ItemGroup {
	var commonChestItems []model.ItemGroup

	commonItems := ir.ItemMap[COMMON]

	for _, item := range commonItems {
		quantity := getRandomNumber(int64(item.MaxQty))
		if quantity == 0 {
			continue
		}
		itemGroup := model.ItemGroup{Item: item, Quantity: quantity}
		commonChestItems = append(commonChestItems, itemGroup)
	}

	return commonChestItems
}

func (ir *ItemRoller) getRareItem(rarity string) model.Item {
	items := ir.ItemMap[rarity]
	return items[getRandomNumber(int64(len(items)))]
}

func (ir *ItemRoller) rollForRareItems() []model.ItemGroup {
	diceRoll := getRandomNumber(100)

	var rareItem model.Item

	switch {
	case diceRoll > 65 && diceRoll <= 85:
		rareItem = ir.getRareItem(RARE)
	case diceRoll > 85 && diceRoll <= 95:
		rareItem = ir.getRareItem(SUPER_RARE)
	case diceRoll > 95 && diceRoll <= 99:
		rareItem = ir.getRareItem(ULTRA_RARE)
	case diceRoll == 100:
		rareItem = ir.getRareItem(ANCIENT)
	}

	if rareItem != (model.Item{}) {
		return []model.ItemGroup{{Item: rareItem, Quantity: 1}}
	}

	return []model.ItemGroup{}
}

func getRandomNumber(max int64) int64 {
	number, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		println(err)
	}

	return number.Int64()
}
