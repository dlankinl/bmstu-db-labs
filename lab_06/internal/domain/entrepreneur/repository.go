package entrepreneur

import (
	"context"
	"database/sql"
	"fmt"
	"lab_02/internal/models"
)

type Repository struct {
	executor *sql.DB
}

func NewEntrepreneurRepository(executor *sql.DB) *Repository {
	return &Repository{executor: executor}
}

func (r Repository) GetEntrepreneurById(ctx context.Context, id int) (ent *Entrepreneur, err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return nil, fmt.Errorf("создание транзакции: %w", err)
	}

	ent = new(Entrepreneur)
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
		`select *
		from lab.enterpreneurs
		where id=$1`,
		id,
	).Scan(
		&ent.Id,
		&ent.FirstName,
		&ent.LastName,
		&ent.Age,
		&ent.Gender,
		&ent.Married,
		&ent.NetWorth,
		&ent.BirthDate,
	)

	if err != nil {
		return nil, fmt.Errorf("поиск предпринимателя по id: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return ent, nil
}

func (r Repository) GetAvgNetWorth(ctx context.Context) (avg float32, err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return 0, fmt.Errorf("создание транзакции: %w", err)
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
		`select avg(net_worth)
		from lab.enterpreneurs`,
	).Scan(&avg)

	if err != nil {
		return 0, fmt.Errorf("нахождение среднего капитала по id: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return avg, nil
}

func (r Repository) GetMaxNetWorth(ctx context.Context) (maxNetWorth float32, err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return 0, fmt.Errorf("создание транзакции: %w", err)
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
		`select get_max_net_worth() as max_net_worth`,
	).Scan(&maxNetWorth)

	if err != nil {
		return 0, fmt.Errorf("вызов скалярной функции из 3-й ЛР: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return maxNetWorth, nil
}

func (r Repository) AddEntrepreneur(ctx context.Context, entrepreneur *models.Entrepreneur) (ent *Entrepreneur, err error) {
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

	var id int
	err = tx.QueryRowContext(
		ctx,
		`select id
    	from insert_entrepreneur(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7)`,
		entrepreneur.FirstName,
		entrepreneur.LastName,
		entrepreneur.Age,
		entrepreneur.Gender,
		entrepreneur.Married,
		entrepreneur.NetWorth,
		entrepreneur.BirthDate,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("вызов табличной функции из 3-й ЛР: %w", err)
	}

	ent = new(Entrepreneur)
	err = tx.QueryRowContext(
		ctx,
		`select * 
		from lab.enterpreneurs
		where id = $1`,
		id,
	).Scan(
		&ent.Id,
		&ent.FirstName,
		&ent.LastName,
		&ent.Age,
		&ent.Gender,
		&ent.Married,
		&ent.NetWorth,
		&ent.BirthDate,
	)
	if err != nil {
		return nil, fmt.Errorf("проверка вставки предпринимателя: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return ent, nil
}

func (r Repository) MarryMalesYoungerThanAge(ctx context.Context, age int) (err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return fmt.Errorf("создание транзакции: %w", err)
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = fmt.Errorf("ошибка: %w, ошибка rollback: %w", err, rollbackErr)
			}
		}
	}()

	_, err = r.executor.ExecContext(
		ctx,
		`call marry_males($1)`,
		age,
	)

	if err != nil {
		return fmt.Errorf("вызов хранимой процедуры из 3-й ЛР: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return nil
}

func (r Repository) CreateTempEntTable(ctx context.Context) (err error) {
	tx, err := r.executor.Begin()
	if err != nil {
		return fmt.Errorf("создание транзакции: %w", err)
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = fmt.Errorf("ошибка: %w, ошибка rollback: %w", err, rollbackErr)
			}
		}
	}()

	_, err = r.executor.ExecContext(
		ctx,
		`create table if not exists lab.entrepreneurs_temp(
		id serial primary key ,
		first_name text not null ,
		last_name text not null,
		age int not null,
		gender boolean not null,
		married boolean not null,
		net_worth int not null,
		birth_date date)`,
	)

	if err != nil {
		return fmt.Errorf("создание временной таблицы предпринимателей: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return nil
}

func (r Repository) AddEntrepreneurTmp(ctx context.Context, entrepreneur *models.Entrepreneur) (ent *Entrepreneur, err error) {
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

	var id int
	err = tx.QueryRowContext(
		ctx,
		`insert into lab.entrepreneurs_temp (
		 	first_name,
		 	last_name,
		 	age,
		 	gender,
		 	married,
		 	net_worth,
		 	birth_date) 
   		values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
        ) returning id`,
		entrepreneur.FirstName,
		entrepreneur.LastName,
		entrepreneur.Age,
		entrepreneur.Gender,
		entrepreneur.Married,
		entrepreneur.NetWorth,
		entrepreneur.BirthDate,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("вызов табличной функции из 3-й ЛР: %w", err)
	}

	ent = new(Entrepreneur)
	err = tx.QueryRowContext(
		ctx,
		`select * 
		from lab.entrepreneurs_temp
		where id = $1`,
		id,
	).Scan(
		&ent.Id,
		&ent.FirstName,
		&ent.LastName,
		&ent.Age,
		&ent.Gender,
		&ent.Married,
		&ent.NetWorth,
		&ent.BirthDate,
	)
	if err != nil {
		return nil, fmt.Errorf("проверка вставки предпринимателя: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("подтверждение транзакции: %w", err)
	}

	return ent, nil
}
