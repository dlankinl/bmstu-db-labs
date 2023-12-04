package others

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	executor *sql.DB
}

func NewOthersRepository(executor *sql.DB) *Repository {
	return &Repository{executor: executor}
}

func (r Repository) GetHostAddrPort(ctx context.Context) (host string, err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return "", fmt.Errorf("создание транзакции: %w", err)
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = fmt.Errorf("ошибка: %w, ошибка rollback: %w", err, rollbackErr)
			}
		}
	}()

	err = tx.QueryRowContext(
		ctx,
		`select addr || ':' || port host
		from inet_server_addr() addr,
			inet_server_port() port`,
	).Scan(&host)

	if err != nil {
		return "", fmt.Errorf("вызов скалярной функции из 3-й ЛР: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return host, nil
}
