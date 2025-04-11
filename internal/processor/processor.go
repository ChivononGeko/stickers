package processor

import (
	"encoding/json"
	"log/slog"
	"stickers/internal/config"
	"stickers/internal/models"
	"stickers/internal/storage"
)

const labelPath = "assets/data/label.pdf"

type Processor struct {
	AllowedIDs  map[string]struct{}
	PrinterName string
	SumatraPDF  string
}

func NewProcessor(config *config.Config) *Processor {
	return &Processor{
		AllowedIDs:  config.AllowedIDs,
		PrinterName: config.PrinterName,
		SumatraPDF:  config.SumatraPDF,
	}
}

func (p *Processor) ProcessOrder(order models.Order) {
	slog.Info("Пришел новый заказ", "client", order.ClientName, "close", order.CloseTime)

	orderNumber := GetNextOrderNumber()

	for _, orderProduct := range order.Products {
		product, found := storage.GetProduct(orderProduct.ID)
		if !found {
			slog.Warn("Товар не найден", "ID", orderProduct.ID)
			continue
		}

		if product.Workshop == "2" {
			continue
		}
		if _, skip := p.AllowedIDs[product.ID]; skip {
			continue
		}

		var mods []models.OrderModification
		if err := json.Unmarshal([]byte(orderProduct.Modifications), &mods); err != nil {
			slog.Warn("Ошибка модификаций", "raw", orderProduct.Modifications, "error", err)
			continue
		}

		var modNames string
		for _, mod := range mods {
			if name := findModifierName(product, mod.M); name != "" {
				if modNames != "" {
					modNames += ", "
				}
				modNames += name
			}
		}

		GeneratePDF(orderNumber, order.ClientName, order.CloseTime, product.Name, modNames, orderProduct.Comment, labelPath)
		PrintPDF(labelPath, p.PrinterName, p.SumatraPDF)
	}
}

func findModifierName(product models.Product, modID int) string {
	for _, group := range product.Modifiers {
		for _, mod := range group.Modifications {
			if mod.ID == modID {
				return mod.Name
			}
		}
	}
	return ""
}
