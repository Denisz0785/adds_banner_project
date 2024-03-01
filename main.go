package main

import "fmt"

// Product Товары выдачи каталога интернет магазина
type Product struct {
	Sku                 string  // артикул товара
	PromotionProviderID *string // Рекламный идентификатор товара
	Price               int     // Цена
}

// AdItem Рекламные или рекламируемые товары
type AdItem struct {
	Sku                 string // артикул товара
	PromotionProviderID string // Рекламный идентификатор товара
}

var SlotPositions = []int{1, 2} // позиции рекламных слотов

// ReplaceSlot устанавливает ads в рекламные слоты, меняя местами с products
func ReplaceSlot(i, y int, ads AdItem, prod *Product, products []*Product, t bool) []*Product {
	//если ads отсутствует в срезе products
	if !t {
		tempID := ads.PromotionProviderID
		temp := products[SlotPositions[i]]
		temp1 := *temp
		temp2 := &temp1
		products = append(products, temp2)
		products[SlotPositions[i]].Sku = ads.Sku
		products[SlotPositions[i]].Price = 0
		products[SlotPositions[i]].PromotionProviderID = &tempID
	} else {
		//если ads присутствует в списке products, то меняем местами
		tempID := ads.PromotionProviderID
		//меняем местами товар из ads и products
		temp := products[SlotPositions[i]]
		products[SlotPositions[i]] = prod
		products[SlotPositions[i]].PromotionProviderID = &tempID
		products[y] = temp
	}
	return products
}

func EnrichProductsWithAds(products []*Product, ads []AdItem) []*Product {
	//инициализация кол-ва рекламных слотов: если >=2, то 2. Если <2, то 1.
	countAds := len(SlotPositions)
	if countAds >= 2 {
		countAds = 2
	} else {
		countAds = 1
	}
	//установка рекламных товаров в нужные слоты
	for i, v := range ads[:countAds] {
		//t определяет были совпадения элемента из ads в products
		t := false
		for y, z := range products {
			if i < countAds && v.Sku == z.Sku {
				t = true
				products = ReplaceSlot(i, y, v, z, products, t)
				break
			} else if !t && y < (len(products)-1) {
				continue
			} else if !t && y == (len(products)-1) {
				a1 := &Product{
					Sku:                 v.Sku,
					PromotionProviderID: &v.PromotionProviderID,
				}
				products = ReplaceSlot(i, y, v, a1, products, t)
			}
		}
	}
	return products
}

var products = []*Product{
	&Product{
		Sku:   "sku1",
		Price: 1,
	},
	&Product{
		Sku:   "sku2",
		Price: 3,
	},
	&Product{
		Sku:   "sku3",
		Price: 1,
	},
	&Product{
		Sku:   "sku4",
		Price: 4,
	},
	&Product{
		Sku:   "sku5",
		Price: 5,
	},
}

var ads = []AdItem{
	{
		Sku:                 "sku4",
		PromotionProviderID: "id1",
	},
	{
		Sku:                 "sku6",
		PromotionProviderID: "id2",
	},
}

func main() {
	newProducts := EnrichProductsWithAds(products, ads)
	for _, v := range newProducts {
		if v.PromotionProviderID != nil {
			fmt.Println(v.Sku, v.Price, *v.PromotionProviderID)
		} else {
			fmt.Println(v.Sku, v.Price)
		}
	}

}
