package company

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	executor *sql.DB
}

func NewCompanyRepository(executor *sql.DB) *Repository {
	return &Repository{executor: executor}
}

func (r Repository) GetCompaniesInCity(ctx context.Context, city string) (companies []Company, err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return nil, fmt.Errorf("создание транзакции: %w", err)
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = fmt.Errorf("ошибка: %w, ошибка rollback: %w", err, rollbackErr)
			}
		}
	}()

	rows, err := tx.QueryContext(
		ctx,
		`select comp.id, comp.name as name, c.name as city, e.name as owner, f.profit as profit
		from lab.companies comp
			join (select id, name
				  from lab.cities
				  where name = $1) c on comp.city_id = c.id
			join (select id, first_name || ' ' || last_name as name
				  from lab.enterpreneurs) e on comp.owner_id = e.id
			join (select id, profit
				  from lab.financials) f on comp.financials_id = f.id`,
		city,
	)

	companies = make([]Company, 0)
	for rows.Next() {
		var tmp Company
		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.City,
			&tmp.Owner,
			&tmp.Profit,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование строки: %w", err)
		}

		companies = append(companies, tmp)
	}

	if err != nil {
		return nil, fmt.Errorf("поиск компаний в городе: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return companies, nil
}

func (r Repository) GetBestCompanies(ctx context.Context, revenue, taxes float32) (companies []CompanyFinancials, err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return nil, fmt.Errorf("создание транзакции: %w", err)
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = fmt.Errorf("ошибка: %w, ошибка rollback: %w", err, rollbackErr)
			}
		}
	}()

	rows, err := tx.QueryContext(
		ctx,
		`with BestCompanies (id, name, owner_id, financials_id) as (
		select id, name, owner_id, financials_id
		from lab.companies
		group by id
		)
		select name, revenue, max(profit) over (partition by name) as Profit
		from lab.financials f join BestCompanies bc
		on f.id = bc.financials_id
		where revenue > $1 and taxes < $2`,
		revenue,
		taxes,
	)

	companies = make([]CompanyFinancials, 0)
	for rows.Next() {
		var tmp CompanyFinancials
		err = rows.Scan(
			&tmp.Name,
			&tmp.Revenue,
			&tmp.Profit,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование строки: %w", err)
		}

		companies = append(companies, tmp)
	}

	if err != nil {
		return nil, fmt.Errorf("поиск лучших компаний: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return companies, nil
}
