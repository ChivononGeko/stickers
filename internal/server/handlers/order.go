package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"stickers/internal/models"
	"stickers/internal/processor"
)

type OrderHandler struct {
	Processor *processor.Processor
}

func NewOrderHandler(p *processor.Processor) *OrderHandler {
	return &OrderHandler{
		Processor: p,
	}
}

func (h *OrderHandler) HandleOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		slog.Warn("Ошибка парсинга JSON", "error", err)
		http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	h.Processor.ProcessOrder(order)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Заказ принят"})
}
