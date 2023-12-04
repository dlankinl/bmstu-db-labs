package meta_data

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	executor *sql.DB
}

func NewMetaDataRepository(executor *sql.DB) *Repository {
	return &Repository{executor: executor}
}

func (r Repository) GetTables(ctx context.Context) (tables []Table, err error) {
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

	rows, err := r.executor.QueryContext(
		ctx,
		`select 
			schemaname, 
			tablename,
			tablename
		from pg_tables`)

	if err != nil {
		return nil, fmt.Errorf("вызов функции с доступом к метаданным: %w", err)
	}

	tables = make([]Table, 0)
	for rows.Next() {
		var tmp Table
		err = rows.Scan(
			&tmp.SchemaName,
			&tmp.Name,
			&tmp.Owner,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование строки: %w", err)
		}

		tables = append(tables, tmp)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return tables, nil
}
