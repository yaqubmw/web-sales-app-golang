package usecase

import (
	"time"

	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/repository"
)

type CompleteReport struct {
	Name      string
	Email     string
	StartDate time.Time
	EndDate   time.Time
	Report    []model.Report
}

type ReportUsecase interface {
	GetReport(startDate, endDate time.Time, requestorName, requestorEmail string) (*CompleteReport, error)
}

type reportUsecase struct {
	reportRepo repository.ReportRepo
}

func (r *reportUsecase) GetReport(startDate, endDate time.Time, requestorName, requestorEmail string) (*CompleteReport, error) {
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.Local)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, time.Local)

	ReportSale, err := r.reportRepo.GetReport(startDate, endDate)
	if err != nil {
		return nil, err
	}
	report := &CompleteReport{
		Name:      requestorName,
		Email:     requestorEmail,
		StartDate: startDate,
		EndDate:   endDate,
		Report:    ReportSale,
	}
	return report, nil
}

func NewReportUsecase(reportRepo repository.ReportRepo) ReportUsecase {
	return &reportUsecase{
		reportRepo: reportRepo,
	}
}
