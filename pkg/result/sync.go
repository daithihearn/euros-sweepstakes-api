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
	// Get the results from the Google Sheet
	res, err := s.SheetService.Srv.Spreadsheets.Values.Get(s.SheetService.SpreadsheetID, "Results!A2:A5").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	if len(res.Values) == 0 {
		return fmt.Errorf("no data found in the sheet")
	}

	resVals := res.Values

	// Get the odds from the Google Sheet
	odds, err := s.SheetService.Srv.Spreadsheets.Values.Get(s.SheetService.SpreadsheetID, "Results!C2:C5").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	if len(odds.Values) == 0 {
		return fmt.Errorf("no data found in the sheet")
	}

	oddsVals := odds.Values

	if len(resVals) < 4 || len(oddsVals) < 4 {
		return fmt.Errorf("invalid data found in the sheet")
	}

	winner, err := parseTeamFromSheet(resVals, oddsVals, 0)
	if err != nil {
		return err
	}
	runnerUp, err := parseTeamFromSheet(resVals, oddsVals, 1)
	if err != nil {
		return err
	}
	thirdPlace, err := parseTeamFromSheet(resVals, oddsVals, 2)
	if err != nil {
		return err
	}
	fourthPlace, err := parseTeamFromSheet(resVals, oddsVals, 3)
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

func parseTeamFromSheet(resVals [][]interface{}, oddsVals [][]interface{}, index int) (team.Team, error) {
	if len(resVals) <= index || len(resVals[index]) == 0 {
		return team.Team{}, fmt.Errorf("no data found")
	}
	countryStr := fmt.Sprintf("%v", resVals[index][0])
	oddsStr := fmt.Sprintf("%v", oddsVals[index][0])
	parsedTeam, err := team.ParseTeam(countryStr, oddsStr)
	if err != nil {
		return team.Team{}, fmt.Errorf("failed to parse: %s %w", countryStr, err)
	}
	return parsedTeam, nil
}
