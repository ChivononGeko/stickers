package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"stickers/internal/models"
	"time"
)

func FetchProducts(apiToken string) ([]models.Product, error) {
	slog.Info("Запрашиваем данные с API...")
	apiURL := fmt.Sprintf("https://joinposter.com/api/menu.getProducts?token=" + apiToken)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse models.APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Response, nil
}
