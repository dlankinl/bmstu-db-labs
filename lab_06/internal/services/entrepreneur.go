package services

import (
	"context"
	"fmt"
	"lab_02/internal/domain/entrepreneur"
	"lab_02/internal/models"
	"time"
)

type EntrepreneurService struct {
	repo entrepreneur.Repository
}

func NewEntrepreneurService(repo entrepreneur.Repository) *EntrepreneurService {
	return &EntrepreneurService{repo: repo}
}

func (s *EntrepreneurService) GetEntrepreneurById(id int) (ent *entrepreneur.Entrepreneur, err error) {
	ctx := context.Background()

	ent, err = s.repo.GetEntrepreneurById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("сервис предпринимателей: %w", err)
	}

	return ent, nil
}

func (s *EntrepreneurService) GetAvgNetWorth() (avg float32, err error) {
	ctx := context.Background()

	avg, err = s.repo.GetAvgNetWorth(ctx)
	if err != nil {
		return 0, fmt.Errorf("сервис предпринимателей: %w", err)
	}

	return avg, nil
}

func (s *EntrepreneurService) GetMaxNetWorth() (maxNetWorth float32, err error) {
	ctx := context.Background()

	maxNetWorth, err = s.repo.GetMaxNetWorth(ctx)
	if err != nil {
		return 0, fmt.Errorf("сервис предпринимателей: %w", err)
	}

	return maxNetWorth, nil
}

func (s *EntrepreneurService) AddEntrepreneur(
	firstName,
	lastName string,
	age int,
	gender,
	married bool,
	netWorth float32,
	birthDate time.Time,
) (err error) {
	ctx := context.Background()

	entrepreneur := models.NewEntrepreneur(firstName, lastName, age, gender, married, netWorth, birthDate)

	ent, err := s.repo.AddEntrepreneur(ctx, entrepreneur)
	if err != nil || ent == nil {
		return fmt.Errorf("создание записи о предпринимателе: %w", err)
	}

	return nil
}

func (s *EntrepreneurService) MarryMalesYoungerThanAge(age int) (err error) {
	ctx := context.Background()

	err = s.repo.MarryMalesYoungerThanAge(ctx, age)
	if err != nil {
		return fmt.Errorf("женитьба мужчин, младше %d: %w", age, err)
	}

	return nil
}

func (s *EntrepreneurService) CreateTempTable() (err error) {
	ctx := context.Background()

	err = s.repo.CreateTempEntTable(ctx)
	if err != nil {
		return fmt.Errorf("создание временной таблицы: %w", err)
	}

	return nil
}

func (s *EntrepreneurService) AddEntrepreneurTmp(
	firstName,
	lastName string,
	age int,
	gender,
	married bool,
	netWorth float32,
	birthDate time.Time,
) (err error) {
	ctx := context.Background()

	entrepreneur := models.NewEntrepreneur(firstName, lastName, age, gender, married, netWorth, birthDate)

	ent, err := s.repo.AddEntrepreneurTmp(ctx, entrepreneur)
	if err != nil || ent == nil {
		return fmt.Errorf("создание записи о предпринимателе во временной таблице: %w", err)
	}

	return nil
}
