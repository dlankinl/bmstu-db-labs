package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"lab_02/internal/utils"
	"os"
	"os/exec"
	"path/filepath"
)

var db *sql.DB

func copyFilesToDocker(dockerID, hostPath, dockerPath string) {
	union := dockerID + ":" + dockerPath
	command := exec.Command("docker", "cp", hostPath, union)

	_, err := command.Output()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=postgres password=postgres host=localhost port=5438 sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
}

func handleQueries(containerID, hostPath, dockerPath string) {
	createQueryPath := filepath.Join("lab_01", "sql", "create.sql")
	query, err := os.ReadFile(createQueryPath)
	if err != nil {
		fmt.Println("error", err)
		db.Close()
		return
	}

	_, err = db.Exec(string(query))
	if err != nil {
		fmt.Println("error", err)
		db.Close()
		return
	}

	createQueryPath = filepath.Join("lab_01", "sql", "constraints.sql")
	query, err = os.ReadFile(createQueryPath)
	if err != nil {
		fmt.Println("error", err)
		db.Close()
		return
	}

	_, err = db.Exec(string(query))
	if err != nil {
		fmt.Println("error", err)
		db.Close()
		return
	}

	copyFilesToDocker(containerID, hostPath, dockerPath)

	createQueryPath = filepath.Join("lab_01", "sql", "copy.sql")
	query, err = os.ReadFile(createQueryPath)
	if err != nil {
		fmt.Println("error", err)
		db.Close()
	}

	_, err = db.Exec(string(query))
	if err != nil {
		fmt.Println("error", err)
		db.Close()
	}
}

func main() {
	var count int
	fmt.Printf("Enter a number of companies to generate: ")
	fmt.Scanln(&count)

	// Paste your values
	containerID := "fed0ec6f6325"
	hostPath := "/Users/dmitry/Desktop/bmstu/db1/lab_01/data/."
	dockerPath := "/var/lib/postgresql/data/csv"

	people, pHeader := utils.GenerateEnterpreneurs(count)
	cities, cHeader := utils.GenerateCities(count)
	financials, fHeader := utils.GenerateFinancials(count)
	companies, coHeader := utils.GenerateCompanies(count)
	skillsDescr, sDHeader := utils.GenerateSkillsDescriptions(count)
	skillsNames, sNHeader := utils.GenerateSkillsNames(count)

	dataPath := filepath.Join("lab_01", "data")

	enterpreneursPath := filepath.Join(dataPath, "enterpreneurs.csv")
	citiesPath := filepath.Join(dataPath, "cities.csv")
	financialsPath := filepath.Join(dataPath, "financials.csv")
	companiesPath := filepath.Join(dataPath, "companies.csv")
	skillsDescrPath := filepath.Join(dataPath, "skills_descr.csv")
	skillsNamePath := filepath.Join(dataPath, "skills_name.csv")

	utils.WriteFile(enterpreneursPath, people, pHeader)
	utils.WriteFile(citiesPath, cities, cHeader)
	utils.WriteFile(financialsPath, financials, fHeader)
	utils.WriteFile(companiesPath, companies, coHeader)
	utils.WriteFile(skillsDescrPath, skillsDescr, sDHeader)
	utils.WriteFile(skillsNamePath, skillsNames, sNHeader)

	handleQueries(containerID, hostPath, dockerPath)

	defer db.Close()
}
