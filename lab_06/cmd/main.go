package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"lab_02/internal/domain/company"
	"lab_02/internal/domain/entrepreneur"
	"lab_02/internal/domain/meta_data"
	"lab_02/internal/domain/others"
	"lab_02/internal/services"
	"log"
	"time"
)

const commands = `
	1. Выполнить скалярный запрос;
	2. Выполнить запрос с несколькими соединениями (JOIN);
	3. Выполнить запрос с ОТВ(CTE) и оконными функциями;
	4. Выполнить запрос к метаданным;
	5. Вызвать скалярную функцию (написанную в третьей лабораторной работе);
	6. Вызвать многооператорную или табличную функцию (написанную в третьей лабораторной работе);
	7. Вызвать хранимую процедуру (написанную в третьей лабораторной работе);
	8. Вызвать системную функцию или процедуру;
	9. Создать таблицу в базе данных, соответствующую тематике БД;
	10. Выполнить вставку данных в созданную таблицу с использованием инструкции INSERT или COPY.

	0. Завершить работу.
	`

func newConnection(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("подключение к БД: %w", err)
	}
	if err = db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("пинг БД: %w", err)
	}

	return db, nil
}

func main() {
	connStr := "postgresql://postgres:postgres@localhost:5438/postgres?sslmode=disable"
	db, err := newConnection(connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	compRepo := company.NewCompanyRepository(db)
	compSvc := services.NewCompanyService(*compRepo)

	entRepo := entrepreneur.NewEntrepreneurRepository(db)
	entSvc := services.NewEntrepreneurService(*entRepo)

	othRepo := others.NewOthersRepository(db)
	othSvc := services.NewOthersService(*othRepo)

	metaRepo := meta_data.NewMetaDataRepository(db)
	metaSvc := services.NewMetaDataService(*metaRepo)

	_ = entSvc

	var cmd int
commandChoice:
	for {
		fmt.Println(commands)

		_, err = fmt.Scanf("%d", &cmd)
		if err != nil {
			log.Fatal(err)
		}

		switch cmd {
		case 0:
			break commandChoice
		case 1:
			avg, err := entSvc.GetAvgNetWorth()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(avg)
		case 2:
			var city string
			fmt.Println("Введите название города: ")
			_, err := fmt.Scanf("%s", &city)
			if err != nil {
				log.Fatal(err)
			}
			companies, err := compSvc.GetCompaniesInCity(city)
			if err != nil {
				log.Fatal(err)
			}
			company.PrintCompanies(companies)
		case 3:
			var revenue, taxes float32
			fmt.Println("Введите размер выручки: ")
			_, err = fmt.Scanf("%f", &revenue)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Введите сумму налогов: ")
			_, err = fmt.Scanf("%f", &taxes)
			if err != nil {
				log.Fatal(err)
			}
			companies, err := compSvc.GetBestCompanies(revenue, taxes)
			if err != nil {
				fmt.Printf("%v", err)
				continue
			}
			company.PrintCompanyFinancials(companies)
		case 4:
			tables, err := metaSvc.GetTables()
			if err != nil {
				fmt.Printf("Получение метаданных: %v", err)
				continue
			}
			meta_data.PrintMetaTables(tables)
		case 5:
			maxNetWorth, err := entSvc.GetMaxNetWorth()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Максимальный капитал: ", maxNetWorth)
		case 6:
			err = addEntrepreneur(entSvc.AddEntrepreneur)
			if err != nil {
				fmt.Printf("%v", err)
				continue
			}
		case 7:
			var age int
			fmt.Printf("Введите возраст: ")
			_, err = fmt.Scanf("%d", &age)
			if err != nil {
				log.Fatal(err)
			}
			err = entSvc.MarryMalesYoungerThanAge(age)
			if err != nil {
				fmt.Printf("Женитьба мужчин: %v", err)
				continue
			}
		case 8:
			host, err := othSvc.GetHostAddrPort()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Хост: ", host)
		case 9:
			err = entSvc.CreateTempTable()
			if err != nil {
				fmt.Printf("Создание таблицы: %v", err)
				continue
			}
		case 10:
			err = addEntrepreneur(entSvc.AddEntrepreneurTmp)
			if err != nil {
				fmt.Printf("%v", err)
				continue
			}
		default:
			continue
		}
	}
}

func addEntrepreneur(addFn func(
	firstName, lastName string,
	age int,
	gender, married bool,
	netWorth float32,
	birthDate time.Time) error) (err error) {
	var firstName, lastName, bDate string
	var age, gndr, mrrd int
	var netWorth float32
	var gender, married bool
	var birthDate time.Time
	fmt.Printf("Введите имя: ")
	_, err = fmt.Scanf("%s", &firstName)
	if err != nil {
		return fmt.Errorf("ошибка ввода имени: %v", err)
	}
	fmt.Printf("Введите фамилию: ")
	_, err = fmt.Scanf("%s", &lastName)
	if err != nil {
		return fmt.Errorf("ошибка ввода фамилии: %v", err)
	}
	fmt.Printf("Введите возраст: ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		return fmt.Errorf("ошибка ввода возраста: %v", err)
	}
	fmt.Printf("Введите пол (0 - женский, 1 - мужской): ")
	_, err = fmt.Scanf("%d", &gndr)
	if err != nil {
		return fmt.Errorf("ошибка ввода пола: %v", err)
	}
	if gndr > 1 || gndr < 0 {
		return fmt.Errorf("пол должен быть либо 0, либо 1")
	}
	if gndr == 1 {
		gender = true
	} else if gndr == 0 {
		gender = false
	}
	fmt.Printf("Введите семейное положение (0 - холост[-а], 1 - женат[замужем]): ")
	_, err = fmt.Scanf("%d", &mrrd)
	if err != nil {
		return fmt.Errorf("ошибка ввода семейного положения: %v", err)
	}
	if mrrd > 1 || mrrd < 0 {
		return fmt.Errorf("семейное положение должно быть либо 0, либо 1")
	}
	if mrrd == 1 {
		married = true
	} else if mrrd == 0 {
		married = false
	}
	fmt.Printf("Введите капитал: ")
	_, err = fmt.Scanf("%f", &netWorth)
	if err != nil {
		return fmt.Errorf("ошибка ввода капитала: %v", err)
	}
	fmt.Printf("Введите дату рождения (формат 'гггг-мм-дд'): ")
	_, err = fmt.Scanf("%s", &bDate)
	if err != nil {
		return fmt.Errorf("ошибка ввода даты рождения: %v", err)
	}
	birthDate, err = time.Parse("2006-01-02", bDate)
	if err != nil {
		return fmt.Errorf("ошибка чтения даты: %v", err)
	}
	return addFn(firstName, lastName, age, gender, married, netWorth, birthDate)
}
