package processor

import (
	"log/slog"

	"github.com/phpdave11/gofpdf"
)

func GeneratePDF(orderNumber, clientName, orderTime, productName, modificatorName, productComment, labelPath string) {
	if clientName == "" {
		clientName = "dosym"
	}

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 58, Ht: 40},
	})
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPage()

	pdf.AddUTF8Font("BebasNeue", "", "assets/font/BebasNeue-Regular.ttf")

	// Верхняя строка:
	pdf.SetFont("BebasNeue", "", 14)
	pdf.SetXY(2, 4)
	pdf.CellFormat(36, 6, "Salem, ", "", 0, "L", false, 0, "")
	pdf.SetFont("BebasNeue", "", 14)
	pdf.SetXY(13, 4)
	pdf.CellFormat(20, 6, clientName, "", 0, "L", false, 0, "")

	// Логотип
	pdf.Image("assets/images/logo.png", 44, 2, 13, 6, false, "", 0, "")

	// Номер заказа
	pdf.SetFont("BebasNeue", "", 30)
	pdf.SetXY(2, 12)
	pdf.CellFormat(54, 8, orderNumber, "", 0, "L", false, 0, "")

	// Название напитка
	pdf.SetFont("BebasNeue", "", 15)
	pdf.SetXY(16, 10)
	pdf.MultiCell(54, 8, productName, "", "L", false)

	// Состав напитка
	pdf.SetFont("BebasNeue", "", 8)
	pdf.SetXY(16, 15)
	pdf.MultiCell(54, 8, modificatorName, "", "L", false)

	// Комментарий товара
	pdf.SetFont("BebasNeue", "", 8)
	pdf.SetXY(16, 19)
	pdf.MultiCell(54, 8, productComment, "", "L", false)

	// Время заказа
	pdf.SetFont("BebasNeue", "", 7)
	pdf.SetXY(2, 31)
	pdf.Cell(54, 6, orderTime)

	err := pdf.OutputFileAndClose(labelPath)
	if err != nil {
		slog.Error("Не удалось сохранить стикер", "error", err)
		return
	}

	slog.Info("Стикер сформирован и сохранен")
}
