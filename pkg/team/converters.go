package team

import (
	"fmt"
	"strings"
)

// ParseTeamFromString the team object from the country string
// the format of the string is "🇪🇸 Spain"
func ParseTeamFromString(input string) (Team, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return Team{}, fmt.Errorf("empty team")
	}

	return Team{
		Country: input,
	}, nil
}
