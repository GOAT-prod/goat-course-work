package service

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"report-service/database"
	"report-service/domain"
	"report-service/repository"
	"time"
)

type Report interface {
	GetFactoryReport(ctx goatcontext.Context, factoryId int, date time.Time) (*excelize.File, error)
	GetUserReport(ctx goatcontext.Context, userId int, date time.Time) (*excelize.File, error)
}

type ReportService struct {
	reportRepository repository.Report
}

func NewReportService(reportRepository repository.Report) Report {
	return &ReportService{
		reportRepository: reportRepository,
	}
}

func (s *ReportService) GetFactoryReport(ctx goatcontext.Context, factoryId int, date time.Time) (*excelize.File, error) {
	dbReportItems, err := s.reportRepository.GetFactoryReportItems(ctx, factoryId, date, date.Add(24*time.Hour))
	if err != nil {
		return nil, err
	}

	itemsPrices := lo.Map(dbReportItems, func(item database.ReportItem, _ int) int {
		return int(item.Price.IntPart())
	})

	return BuildReportXLSX(domain.Report{
		Date:  date,
		Total: decimal.NewFromInt32(int32(lo.Sum(itemsPrices))),
		Items: lo.Map(dbReportItems, func(item database.ReportItem, _ int) domain.ReportItem {
			return domain.ReportItem{
				ProductName: item.ProductName,
				Color:       item.Color,
				Size:        item.Size,
				Count:       item.Count,
				Price:       item.Price,
			}
		}),
	})
}

func (s *ReportService) GetUserReport(ctx goatcontext.Context, userId int, date time.Time) (*excelize.File, error) {
	dbReportItems, err := s.reportRepository.GetUserReportItems(ctx, userId, date, date.Add(24*time.Hour))
	if err != nil {
		return nil, err
	}

	itemsPrices := lo.Map(dbReportItems, func(item database.ReportItem, _ int) int {
		return int(item.Price.IntPart())
	})

	return BuildReportXLSX(domain.Report{
		Date:  date,
		Total: decimal.NewFromInt32(int32(lo.Sum(itemsPrices))),
		Items: lo.Map(dbReportItems, func(item database.ReportItem, _ int) domain.ReportItem {
			return domain.ReportItem{
				ProductName: item.ProductName,
				Color:       item.Color,
				Size:        item.Size,
				Count:       item.Count,
				Price:       item.Price,
			}
		}),
	})
}

func BuildReportXLSX(report domain.Report) (*excelize.File, error) {
	f := excelize.NewFile()
	sheetName := "Отчет"

	// Переименовываем лист по умолчанию
	defaultSheet := f.GetSheetName(0)
	err := f.SetSheetName(defaultSheet, sheetName)
	if err != nil {
		return nil, err
	}

	// Заполняем общую информацию по отчёту
	err = f.SetCellValue(sheetName, "A1", "Дата")
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "B1", report.Date.Format("02.01.2006"))
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "D1", "Итого")
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "E1", report.Total.StringFixed(2))
	if err != nil {
		return nil, err
	}

	// Заголовки таблицы с товарами
	err = f.SetCellValue(sheetName, "A3", "Название товара")
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "B3", "Цвет")
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "C3", "Размер")
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "D3", "Количество")
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "E3", "Цена")
	if err != nil {
		return nil, err
	}

	rowStart := 4
	for i, item := range report.Items {
		row := rowStart + i
		err = f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), item.ProductName)
		if err != nil {
			return nil, err
		}

		err = f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), item.Color)
		if err != nil {
			return nil, err
		}

		err = f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), item.Size)
		if err != nil {
			return nil, err
		}

		err = f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), item.Count)
		if err != nil {
			return nil, err
		}

		err = f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), item.Price.StringFixed(2))
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}
