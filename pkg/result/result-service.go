package result

import (
	"encoding/json"
	"euros-sweepstakes-api/pkg/cache"
	"euros-sweepstakes-api/pkg/sheets"
	"euros-sweepstakes-api/pkg/team"
	"fmt"
	"log"
)

type ServiceI interface {
	GetResults() (Result, error)
	RefreshResults() error
}

type Service struct {
	Cache        cache.Cache[Result]
	SheetService *sheets.SheetService
}

func (s *Service) GetResults() (Result, error) {
	result, found, err := s.Cache.Get("results")
	if err != nil {
		return Result{}, err
	}
	if found {
		return result, nil
	}
	// Return an empty struct if the results are not found in the cache
	return Result{}, nil
}

func (s *Service) RefreshResults() error {

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

	winner, err := team.ParseTeamFromString(fmt.Sprintf("%v", val[0]))
	if err != nil {
		return fmt.Errorf("failed to parse winner: %s %w", val[0], err)
	}
	runnerUp, err := team.ParseTeamFromString(fmt.Sprintf("%v", val[1]))
	if err != nil {
		return fmt.Errorf("failed to parse runner up: %s %w", val[1], err)
	}
	thirdPlace, err := team.ParseTeamFromString(fmt.Sprintf("%v", val[2]))
	if err != nil {
		return fmt.Errorf("failed to parse third place: %s %w", val[2], err)
	}
	fourthPlace, err := team.ParseTeamFromString(fmt.Sprintf("%v", val[3]))
	if err != nil {
		return fmt.Errorf("failed to parse fourth place: %s %w", val[3], err)
	}

	result := Result{
		Winner:      winner,
		RunnerUp:    runnerUp,
		ThirdPlace:  thirdPlace,
		FourthPlace: fourthPlace,
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Error marshaling struct to JSON: %v", err)
	}

	// Log the JSON string
	log.Println(string(jsonData))

	err = s.Cache.Set("results", result, 0)
	if err != nil {
		return fmt.Errorf("failed to set results in cache: %w", err)
	}

	return nil
}
