package sheets

import (
	"euros-sweepstakes-api/pkg/score"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ParseScoreFromEntry(entry Entry) (score.Score, error) {
	// Parse the name from the email
	name, err := ParseNameFromEmail(entry.EmailAddress)
	if err != nil {
		return score.Score{}, fmt.Errorf("failed to parse name from email: %w", err)
	}

	winner, err := ParseTeamFromString(entry.Winner)
	if err != nil {
		return score.Score{}, fmt.Errorf("failed to parse winner: %s %w", entry.Winner, err)
	}
	second, err := ParseTeamFromString(entry.SecondPlace)
	if err != nil {
		return score.Score{}, fmt.Errorf("failed to parse second place: %s %w", entry.SecondPlace, err)
	}
	third, err := ParseTeamFromString(entry.ThirdPlace)
	if err != nil {
		return score.Score{}, fmt.Errorf("failed to parse third place: %s %w", entry.ThirdPlace, err)
	}
	fourth, err := ParseTeamFromString(entry.FourthPlace)
	if err != nil {
		return score.Score{}, fmt.Errorf("failed to parse fourth place: %s %w", entry.FourthPlace, err)
	}

	return score.Score{
		Name:       name,
		Teams:      []score.Team{winner, second, third, fourth},
		TotalScore: 0,
	}, nil

}

func ParseNameFromEmail(email string) (string, error) {
	// Split the email address at the '@' symbol
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid email format")
	}

	// Split the name part at the '.' symbol
	nameParts := strings.Split(parts[0], ".")
	if len(nameParts) < 2 {
		return "", fmt.Errorf("invalid name format in email")
	}

	// Capitalize the first letter of each name part and join them with a space
	var nameBuilder strings.Builder
	for i, part := range nameParts {
		if len(part) == 0 {
			continue
		}
		if i > 0 {
			nameBuilder.WriteRune(' ')
		}
		// Capitalize the first letter and make the rest lowercase
		for j, r := range part {
			if j == 0 {
				nameBuilder.WriteRune(unicode.ToUpper(r))
			} else {
				nameBuilder.WriteRune(unicode.ToLower(r))
			}
		}
	}

	return nameBuilder.String(), nil
}

// ParseTeamFromString the team object from the country string
// the format of the string is "🇪🇸 Spain"
func ParseTeamFromString(input string) (score.Team, error) {
	input = strings.TrimSpace(input)
	if len(input) < 2 {
		return score.Team{}, fmt.Errorf("input too short")
	}

	// Decode the first rune
	r1, size1 := utf8.DecodeRuneInString(input)
	if !unicode.IsSymbol(r1) && !unicode.Is(unicode.Regional_Indicator, r1) && !isEmoji(r1) {
		return score.Team{}, fmt.Errorf("invalid flag emoji")
	}

	// Find the end of the emoji sequence
	pos := size1
	for pos < len(input) {
		r, size := utf8.DecodeRuneInString(input[pos:])
		if !isEmojiTag(r) && !isEmoji(r) {
			break
		}
		pos += size
	}

	flag := input[:pos]
	country := strings.TrimSpace(input[pos:])

	// Check for any unexpected characters (e.g., control characters) and remove them
	country = sanitizeString(country)

	return score.Team{
		Country: country,
		Flag:    flag,
		Score:   0,
	}, nil
}

// isEmoji checks if a rune is an emoji
func isEmoji(r rune) bool {
	return unicode.Is(unicode.Regional_Indicator, r) || unicode.IsSymbol(r) || (r >= 0x1F600 && r <= 0x1F64F) // extend range as needed
}

// isEmojiTag checks if a rune is a tag character used in emoji sequences
func isEmojiTag(r rune) bool {
	return r >= 0xE0060 && r <= 0xE007F
}

// sanitizeString removes unexpected control characters from the string
func sanitizeString(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if unicode.IsPrint(r) && !unicode.IsControl(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
