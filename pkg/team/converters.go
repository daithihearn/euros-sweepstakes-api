package team

import (
	"fmt"
	"strings"
)

// ParseTeam the team object from the country string
// the format of the string is "Spain"
func ParseTeam(country string, odds string) (Team, error) {
	country = strings.TrimSpace(country)

	err := validateCountry(country)
	if err != nil {
		return Team{}, err
	}

	return Team{
		Country: country,
		Odds:    odds,
	}, nil
}

// validateCountry validates the country string
// Should be a valid country EU string without any padding in camel case
func validateCountry(country string) error {
	switch country {
	case "Albania", "Andorra", "Armenia", "Austria", "Azerbaijan", "Belarus", "Belgium", "Bosnia and Herzegovina", "Bulgaria", "Croatia", "Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Georgia", "Germany", "Greece", "Hungary", "Iceland", "Ireland", "Italy", "Kazakhstan", "Kosovo", "Latvia", "Liechtenstein", "Lithuania", "Luxembourg", "Malta", "Moldova", "Monaco", "Montenegro", "Netherlands", "North Macedonia", "Norway", "Poland", "Portugal", "Romania", "Russia", "San Marino", "Serbia", "Slovakia", "Slovenia", "Spain", "Sweden", "Switzerland", "Turkey", "Ukraine", "United Kingdom", "Vatican City", "England", "Scotland", "Wales", "Northern Ireland":
		return nil
	}
	return fmt.Errorf("invalid country")
}
