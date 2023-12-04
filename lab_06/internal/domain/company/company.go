package company

import "fmt"

type Company struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	City   string  `json:"city"`
	Owner  string  `json:"owner"`
	Profit float32 `json:"profit"`
}

func PrintCompanies(companies []Company) {
	for _, comp := range companies {
		fmt.Printf("%d: city='%s', name='%s', owner='%s', profit=%.1f\n", comp.Id, comp.City, comp.Name, comp.Owner, comp.Profit)
	}
}

type CompanyFinancials struct {
	Name    string  `json:"name"`
	Revenue float32 `json:"revenue"`
	Profit  float32 `json:"profit"`
}

func PrintCompanyFinancials(companies []CompanyFinancials) {
	for _, comp := range companies {
		fmt.Printf("name='%s', revenue='%.1f', profit=%.1f\n", comp.Name, comp.Revenue, comp.Profit)
	}
}
