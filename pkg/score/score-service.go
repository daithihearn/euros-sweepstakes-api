package score

import (
	"euros-sweepstakes-api/pkg/cache"
	"euros-sweepstakes-api/pkg/result"
	"euros-sweepstakes-api/pkg/sheets"
	"euros-sweepstakes-api/pkg/team"
	"fmt"
)

type ServiceI interface {
	GetScores() ([]Score, error)
	RefreshScores() error
}

type Service struct {
	Cache         cache.Cache[[]Score]
	ResultService result.ServiceI
	SheetService  *sheets.SheetService
}

func (s *Service) GetScores() ([]Score, error) {
	scores, found, err := s.Cache.Get("scores")
	if err != nil {
		return nil, err
	}
	if found {
		return scores, nil
	}
	// Return an empty slice if the scores are not found in the cache
	return []Score{}, nil
}

func (s *Service) RefreshScores() error {
	resp, err := s.SheetService.Srv.Spreadsheets.Values.Get(s.SheetService.SpreadsheetID, "Form Responses 3!B:G").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	// Get the results from the result service
	actual, err := s.ResultService.GetResults()

	var scores []Score
	for i, row := range resp.Values {
		if i == 0 {
			// Skip the header row
			continue
		}
		if len(row) < 6 {
			continue // Skip rows that do not have all columns
		}
		name := fmt.Sprintf("%v", row[5])

		winner, err := team.ParseTeamFromString(fmt.Sprintf("%v", row[1]))
		if err != nil {
			return fmt.Errorf("failed to parse winner: %s %w", row[1], err)
		}
		runnerUp, err := team.ParseTeamFromString(fmt.Sprintf("%v", row[2]))
		if err != nil {
			return fmt.Errorf("failed to parse runner up: %s %w", row[2], err)
		}
		thirdPlace, err := team.ParseTeamFromString(fmt.Sprintf("%v", row[3]))
		if err != nil {
			return fmt.Errorf("failed to parse third place: %s %w", row[3], err)
		}
		fourthPlace, err := team.ParseTeamFromString(fmt.Sprintf("%v", row[4]))
		if err != nil {
			return fmt.Errorf("failed to parse fourth place: %s %w", row[4], err)
		}

		r := result.Result{Winner: winner, RunnerUp: runnerUp, ThirdPlace: thirdPlace, FourthPlace: fourthPlace}

		totalScore := calculateTotalScore(r, actual)

		score := Score{
			Name:       name,
			Result:     r,
			TotalScore: totalScore,
		}

		scores = append(scores, score)
	}

	// Log the number of scores that were found
	fmt.Printf("Refreshing %d scores\n", len(scores))

	err = s.Cache.Set("scores", scores, 0)
	if err != nil {
		return fmt.Errorf("failed to set scores in cache: %w", err)
	}

	return nil
}

// calculateTotalScore calculates the total score for a user based on the proposed and actual results
// 5 points are awarded if a proposed team is present in the actual results
// A bonus of 3 points is awarded if the proposed team is in the correct position
func calculateTotalScore(proposed result.Result, actual result.Result) int {
	res := 0
	if proposed.Winner.Country == actual.Winner.Country {
		res += 3
	}
	if proposed.RunnerUp.Country == actual.RunnerUp.Country {
		res += 3
	}
	if proposed.ThirdPlace.Country == actual.ThirdPlace.Country {
		res += 3
	}
	if proposed.FourthPlace.Country == actual.FourthPlace.Country {
		res += 3
	}

	for _, t := range []team.Team{proposed.Winner, proposed.RunnerUp, proposed.ThirdPlace, proposed.FourthPlace} {
		if t.Country == actual.Winner.Country || t.Country == actual.RunnerUp.Country || t.Country == actual.ThirdPlace.Country || t.Country == actual.FourthPlace.Country {
			res += 5
		}
	}

	return res
}
