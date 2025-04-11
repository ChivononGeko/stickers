package processor

import (
	"fmt"
	"os"
	"strconv"
)

const orderNumberFile = "assets/data/order_number.txt"

func GetNextOrderNumber() string {
	data, err := os.ReadFile(orderNumberFile)
	if err != nil || len(data) == 0 {
		_ = os.WriteFile(orderNumberFile, []byte("001"), 0644)
		return "001"
	}

	current, err := strconv.Atoi(string(data))
	if err != nil {
		_ = os.WriteFile(orderNumberFile, []byte("001"), 0644)
		return "001"
	}

	next := current + 1
	if next > 999 {
		next = 1
	}
	order := fmt.Sprintf("%03d", next)

	_ = os.WriteFile(orderNumberFile, []byte(order), 0644)
	return order
}
