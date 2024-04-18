package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"github.com/yaqubmw/web-sales-app-golang/delivery/middleware"
	"github.com/yaqubmw/web-sales-app-golang/usecase"
	"github.com/yaqubmw/web-sales-app-golang/utils/security"
)

type ReportController struct {
	router   *gin.Engine
	reportUC usecase.ReportUsecase
}

func (r *ReportController) GetReport(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}

	claims, err := security.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}

	requestorName, ok := claims["Name"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}

	requestorEmail, ok := claims["Email"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}

	startDate, err := time.Parse("2006-01-02", c.Query("start_date"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid start date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", c.Query("end_date"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid end date"})
		return
	}

	report, err := r.reportUC.GetReport(startDate, endDate, requestorName, requestorEmail)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to get report"})
		return
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sales Report")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to add sheet"})
		return
	}

	requestorRow := sheet.AddRow()
	TitleStyle := xlsx.NewStyle()
	TitleStyle.Alignment.Horizontal = "left"
	TitleStyle.Alignment.Vertical = "center"
	TitleStyle.Font.Bold = true
	ValueStyle := xlsx.NewStyle()
	ValueStyle.Alignment.Horizontal = "left"
	ValueStyle.Alignment.Vertical = "center"
	ValueStyle.Font.Bold = false
	ValueGreenStyle := xlsx.NewStyle()
	ValueGreenStyle.Alignment.Horizontal = "left"
	ValueGreenStyle.Alignment.Vertical = "center"
	ValueGreenStyle.Font.Bold = false
	ValueGreenStyle.Font.Color = "00FF00"

	requestorCell := requestorRow.AddCell()
	requestorCell.SetString("Requestor")
	requestorCell.SetStyle(TitleStyle)
	requestorCell = requestorRow.AddCell()
	requestorCell.SetString(requestorName)
	requestorCell.SetStyle(ValueStyle)
	requestorCell = requestorRow.AddCell()
	requestorCell.SetString(requestorEmail)
	requestorCell.SetStyle(ValueGreenStyle)

	blankSpaceRow := sheet.AddRow()
	blankSpaceCell := blankSpaceRow.AddCell()
	blankSpaceCell.SetString("")

	parameterRow := sheet.AddRow()
	parameterCell := parameterRow.AddCell()
	parameterCell.SetString("Parameter")
	parameterCell.SetStyle(TitleStyle)

	startDateRow := sheet.AddRow()
	startDateCell := startDateRow.AddCell()
	startDateCell.SetString("Start Date")
	startDateCell.SetStyle(TitleStyle)
	startDateCell = startDateRow.AddCell()
	startDateCell.SetString(startDate.Format("12 April 2006"))
	startDateCell.SetStyle(ValueStyle)

	endDateRow := sheet.AddRow()
	endDateCell := endDateRow.AddCell()
	endDateCell.SetString("End Date")
	endDateCell.SetStyle(TitleStyle)
	endDateCell = endDateRow.AddCell()
	endDateCell.SetString(endDate.Format("12 April 2006"))
	endDateCell.SetStyle(ValueStyle)

	blankSpaceCell = endDateRow.AddCell()
	blankSpaceCell.SetString("")

	headerRow := sheet.AddRow()
	headerStyle := xlsx.NewStyle()
	headerStyle.Alignment.Horizontal = "center"
	headerStyle.Alignment.Vertical = "center"
	headerStyle.Font.Bold = true

	headerCell := headerRow.AddCell()
	headerCell.SetString("User")
	headerCell.SetStyle(headerStyle)
	headerCell = headerRow.AddCell()
	headerCell.SetString("Jumlah Hari Kerja")
	headerCell.SetStyle(headerStyle)
	headerCell = headerRow.AddCell()
	headerCell.SetString("Jumlah Transaksi Barang")
	headerCell.SetStyle(headerStyle)
	headerCell = headerRow.AddCell()
	headerCell.SetString("Jumlah Transaksi Jasa")
	headerCell.SetStyle(headerStyle)
	headerCell = headerRow.AddCell()
	headerCell.SetString("Nominal Transaksi Barang")
	headerCell.SetStyle(headerStyle)
	headerCell = headerRow.AddCell()
	headerCell.SetString("Nominal Transaksi Jasa")
	headerCell.SetStyle(headerStyle)

	for _, r := range report.Report {
		dataRow := sheet.AddRow()

		dataCell := dataRow.AddCell()
		dataCell.SetString(r.User)
		dataCell.SetStyle(ValueStyle)

		dataCell = dataRow.AddCell()
		dataCell.SetInt(r.JumlahHariKerja)
		dataCell.SetStyle(ValueStyle)

		dataCell = dataRow.AddCell()
		dataCell.SetInt(r.JumlahTransaksiBarang)
		dataCell.SetStyle(ValueStyle)

		dataCell = dataRow.AddCell()
		dataCell.SetInt(r.JumlahTransaksiJasa)
		dataCell.SetStyle(ValueStyle)

		dataCell = dataRow.AddCell()
		dataCell.SetFloat(r.NominalTransaksiBarang)
		dataCell.SetStyle(ValueStyle)

		dataCell = dataRow.AddCell()
		dataCell.SetFloat(r.NominalTransaksiJasa)
		dataCell.SetStyle(ValueStyle)
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=SalesReport+"+requestorName+"-"+startDate.Format("12 April 2006")+"-"+endDate.Format("12 April 2006")+".xlsx")

	err = file.Write(c.Writer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to write file"})
		return
	}
}

func NewReportController(router *gin.Engine, reportUC usecase.ReportUsecase) *ReportController {
	controller := ReportController{
		router:   router,
		reportUC: reportUC,
	}

	rg := router.Group("/api/sales")
	rg.GET("/report", middleware.AuthMiddleware(), controller.GetReport)

	return &controller
}
