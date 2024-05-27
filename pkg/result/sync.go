package result

import (
	"euros-sweepstakes-api/pkg/sheets"
	"euros-sweepstakes-api/pkg/team"
	"fmt"
)

type SyncI interface {
	RefreshResults() error
}

type Sync struct {
	ResultService ServiceI
	SheetService  *sheets.SheetService
}

func (s *Sync) RefreshResults() error {
	resp, err := s.SheetService.Srv.Spreadsheets.Values.Get(s.SheetService.SpreadsheetID, "Results!A2:A5").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return fmt.Errorf("no data found in the sheet")
	}

	val := resp.Values

	if len(val) < 4 {
		return fmt.Errorf("invalid data found in the sheet")
	}

	winner, err := parseTeamFromSheet(val, 0, "winner")
	if err != nil {
		return err
	}
	runnerUp, err := parseTeamFromSheet(val, 1, "runner up")
	if err != nil {
		return err
	}
	thirdPlace, err := parseTeamFromSheet(val, 2, "third place")
	if err != nil {
		return err
	}
	fourthPlace, err := parseTeamFromSheet(val, 3, "fourth place")
	if err != nil {
		return err
	}

	result := Result{
		Winner:      winner,
		RunnerUp:    runnerUp,
		ThirdPlace:  thirdPlace,
		FourthPlace: fourthPlace,
	}

	return s.ResultService.SetResult(result)
}

func parseTeamFromSheet(values [][]interface{}, index int, position string) (team.Team, error) {
	if len(values) <= index || len(values[index]) == 0 {
		return team.Team{}, fmt.Errorf("no data for %s", position)
	}
	teamStr := fmt.Sprintf("%v", values[index][0])
	parsedTeam, err := team.ParseTeamFromString(teamStr)
	if err != nil {
		return team.Team{}, fmt.Errorf("failed to parse %s: %s %w", position, teamStr, err)
	}
	return parsedTeam, nil
}
