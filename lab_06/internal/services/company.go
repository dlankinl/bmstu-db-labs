package services

import (
	"context"
	"fmt"
	"lab_02/internal/domain/company"
)

type CompanyService struct {
	repo company.Repository
}

func NewCompanyService(repo company.Repository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) GetCompaniesInCity(city string) (companies []company.Company, err error) {
	ctx := context.Background()
	cmps, err := s.repo.GetCompaniesInCity(ctx, city)
	if err != nil {
		return nil, fmt.Errorf("cервис компаний: %w", err)
	}

	companies = make([]company.Company, 0)
	for _, comp := range cmps {
		companies = append(companies, comp)
	}

	return companies, nil
}

func (s *CompanyService) GetBestCompanies(revenue, taxes float32) (companies []company.CompanyFinancials, err error) {
	ctx := context.Background()
	cmps, err := s.repo.GetBestCompanies(ctx, revenue, taxes)
	if err != nil {
		return nil, fmt.Errorf("cервис компаний: %w", err)
	}

	companies = make([]company.CompanyFinancials, 0)
	for _, comp := range cmps {
		companies = append(companies, comp)
	}

	return companies, nil
}
