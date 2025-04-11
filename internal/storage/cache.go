package storage

import (
	"encoding/json"
	"log/slog"
	"os"
	"stickers/internal/models"
)

var productData = make(map[string]models.Product)

// SetProductData сохраняет товары в кеш
func SetProductData(products []models.Product) {
	for _, product := range products {
		productData[product.ID] = product
	}

	// Сохраняем кеш в файл
	SaveToFile("cache/products_cache.json", products)
}

// GetProduct возвращает товар по ID
func GetProduct(id string) (models.Product, bool) {
	product, found := productData[id]
	return product, found
}

// SaveToFile сохраняет товары в JSON-файл
func SaveToFile(filename string, data []models.Product) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		slog.Warn("Ошибка при сохранении данных", "error", err)
		return
	}
	os.WriteFile(filename, file, 0644)
}

// LoadFromFile загружает товары из JSON-файла
func LoadFromFile(filename string) ([]models.Product, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	return products, nil
}
