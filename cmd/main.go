package main

import (
	"log"
	"log/slog"
	"net/http"
	"stickers/internal/api"
	"stickers/internal/config"
	"stickers/internal/processor"
	"stickers/internal/server"
	"stickers/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	products, err := api.FetchProducts(cfg.PosterToken)
	if err != nil {
		slog.Warn("Ошибка загрузки API, пробуем кеш...", "error", err)
		products, _ = storage.LoadFromFile("cache/products_cache.json")
	}
	slog.Info("Модификаторы определены")

	storage.SetProductData(products)

	p := processor.NewProcessor(cfg)
	s := server.NewServer(p)

	slog.Info("Сервер запущен", "address", cfg.PortServer)
	err = http.ListenAndServe(":"+cfg.PortServer, s.Router())
	if err != nil {
		slog.Error("Ошибка запуска сервера", "error", err)
	}
}
