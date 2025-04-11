package processor

import (
	"log/slog"
	"os/exec"
)

func PrintPDF(filePath, printer, sumatra string) {
	cmd := exec.Command(sumatra,
		"-print-to", printer,
		"-print-settings", "landscape",
		"-silent",
		filePath)

	if err := cmd.Run(); err != nil {
		slog.Error("Ошибка при печати", "error", err)
		return
	}

	slog.Info("Печать завершена")
}
