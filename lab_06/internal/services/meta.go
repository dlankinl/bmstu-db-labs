package services

import (
	"context"
	"fmt"
	"lab_02/internal/domain/meta_data"
)

type MetaDataService struct {
	repo meta_data.Repository
}

func NewMetaDataService(repo meta_data.Repository) *MetaDataService {
	return &MetaDataService{repo: repo}
}

func (s *MetaDataService) GetTables() (tables []meta_data.Table, err error) {
	ctx := context.Background()

	tbls, err := s.repo.GetTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("cервис метаданных: %w", err)
	}

	tables = make([]meta_data.Table, 0)
	for _, table := range tbls {
		tables = append(tables, table)
	}

	return tables, nil
}
