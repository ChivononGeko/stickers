package routes

import (
	"net/http"
	"stickers/internal/server/handlers"
)

func RegisterRoutes(mux *http.ServeMux, orderHandler *handlers.OrderHandler) {
	mux.HandleFunc("/order", orderHandler.HandleOrder)
}
