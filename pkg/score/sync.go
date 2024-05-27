package score

import (
	"euros-sweepstakes-api/pkg/result"
	"euros-sweepstakes-api/pkg/sheets"
	"euros-sweepstakes-api/pkg/team"
	"fmt"
)

type SyncI interface {
	RefreshScores() error
}

type Sync struct {
	ResultService *result.Service
	ScoreService  *Service
	SheetService  *sheets.SheetService
}

func (s *Sync) RefreshScores() error {
	resp, err := s.SheetService.Srv.Spreadsheets.Values.Get(s.SheetService.SpreadsheetID, "Form Responses 3!B:G").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	// Get the results from the result service
	actual, err := s.ResultService.GetResult()

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

	return s.ScoreService.SetScores(scores)
}
