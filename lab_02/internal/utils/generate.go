package utils

import (
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomIntSlice(ran *rand.Rand, amount int) []int {
	nums := make([]int, amount)

	for i := 0; i < amount; i++ {
		nums[i] = i + 1
	}
	ran.Shuffle(amount, func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	return nums
}

func randomDate() time.Time {
	min_ := time.Date(1933, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max_ := time.Date(2005, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max_ - min_

	sec := rand.Int63n(delta) + min_
	return time.Unix(sec, 0)
}

func generateEnterpreneur(ran *rand.Rand) []string {
	gender := ran.Intn(2)
	firstName := randomdata.FirstName(1 - gender)
	lastName := randomdata.LastName()
	birthDate := randomDate()
	age := strconv.Itoa(int(time.Now().Sub(birthDate).Hours() / 24 / 365))
	isMale := strconv.FormatBool(gender == 1)
	married := strconv.FormatBool(ran.Intn(2) == 1)
	netWorth := strconv.Itoa(ran.Intn(99000000) + 1000000)
	return []string{firstName, lastName, age, isMale, married, netWorth, birthDate.Format("2006-01-02")}
}

func GenerateEnterpreneurs(count int) ([][]string, []string) {
	people := make([][]string, count)

	seed := rand.NewSource(time.Now().Unix())
	ran := rand.New(seed)

	for i := 0; i < count; i++ {
		people[i] = generateEnterpreneur(ran)
	}

	header := []string{"first_name", "last_name", "age", "gender", "married", "net_worth", "birth_date"}

	return people, header
}

func generateCity(ran *rand.Rand) []string {
	name := randomdata.City()
	population := strconv.Itoa(ran.Intn(1000000))
	return []string{name, population}
}

func GenerateCities(count int) ([][]string, []string) {
	cities := make([][]string, count)

	seed := rand.NewSource(time.Now().Unix())
	ran := rand.New(seed)

	for i := 0; i < count; i++ {
		cities[i] = generateCity(ran)
	}

	header := []string{"name", "population"}

	return cities, header
}

func generateFinancial(ran *rand.Rand) []string {
	revenue := strconv.Itoa(ran.Intn(9900000) + 100000)
	profit := ran.Intn(990000) + 10000
	totalAssets := strconv.Itoa(ran.Intn(99000000) + 1000000)
	taxRate := rand.Intn(26) + 5
	taxes := strconv.Itoa(profit * taxRate / 100)
	profitStr := strconv.Itoa(profit)
	return []string{revenue, profitStr, totalAssets, taxes}
}

func GenerateFinancials(count int) ([][]string, []string) {
	financials := make([][]string, count)

	seed := rand.NewSource(time.Now().Unix())
	ran := rand.New(seed)

	for i := 0; i < count; i++ {
		financials[i] = generateFinancial(ran)
	}

	header := []string{"revenue", "profit", "total_assets", "taxes"}

	return financials, header
}

func generateCompany(ownerID, finID, cityID int) []string {
	name := randomdata.SillyName()
	ownerIDStr := strconv.Itoa(ownerID)
	finIDStr := strconv.Itoa(finID)
	cityIDStr := strconv.Itoa(cityID)
	return []string{name, ownerIDStr, cityIDStr, finIDStr}
}

func GenerateCompanies(count int) ([][]string, []string) {
	companies := make([][]string, count)

	seed := rand.NewSource(time.Now().Unix())
	ran := rand.New(seed)

	finIDs := generateRandomIntSlice(ran, count)
	citiesIDs := generateRandomIntSlice(ran, count)
	ownerIDs := generateRandomIntSlice(ran, count)

	for i := 0; i < count; i++ {
		companies[i] = generateCompany(ownerIDs[i], finIDs[i], citiesIDs[i])
	}

	header := []string{"name", "owner_id", "city_id", "fin_id"}

	return companies, header
}

func generateSkillDescription(ran *rand.Rand) []string {
	descr := randomdata.Paragraph()
	return []string{descr}
}

func GenerateSkillsDescriptions(count int) ([][]string, []string) {
	descriptions := make([][]string, count)

	seed := rand.NewSource(time.Now().Unix())
	ran := rand.New(seed)

	for i := 0; i < count; i++ {
		descriptions[i] = generateSkillDescription(ran)
	}

	header := []string{"description"}

	return descriptions, header
}

func generateSkillName(entID, skillID int) []string {
	name := randomdata.SillyName()
	entIDStr := strconv.Itoa(entID)
	skillIDStr := strconv.Itoa(skillID)
	return []string{entIDStr, name, skillIDStr}
}

func GenerateSkillsNames(count int) ([][]string, []string) {
	names := make([][]string, count)

	seed := rand.NewSource(time.Now().Unix())
	ran := rand.New(seed)

	entIDs := generateRandomIntSlice(ran, count)
	skillIDs := generateRandomIntSlice(ran, count)

	for i := 0; i < count; i++ {
		names[i] = generateSkillName(entIDs[i], skillIDs[i])
	}

	header := []string{"ent_id", "name", "skill_id"}

	return names, header
}
