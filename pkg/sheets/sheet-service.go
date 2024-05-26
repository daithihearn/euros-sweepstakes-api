package sheets

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
)

type SheetService struct {
	srv           *sheets.Service
	spreadsheetID string
	readRange     string
}

func NewSheetService(ctx context.Context) (*SheetService, error) {

	// Get the service account JSON from the environment variable
	saJSON := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	if saJSON == "" {
		return nil, fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS_JSON environment variable not set")
	}

	// Get the spreadsheet ID and range from environment variables
	spreadsheetID := os.Getenv("GOOGLE_SPREADSHEET_ID")
	if spreadsheetID == "" {
		return nil, fmt.Errorf("GOOGLE_SPREADSHEET_ID environment variable not set")
	}

	// Create a new Sheets service with the service account
	srv, err := sheets.NewService(ctx, option.WithCredentialsJSON([]byte(saJSON)))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	return &SheetService{
		srv:           srv,
		spreadsheetID: spreadsheetID,
		readRange:     "Form Responses 3!A:F",
	}, nil
}

// GetEntries retrieves the data from the Google Sheet and returns a slice of Entry structs
func (s *SheetService) GetEntries() ([]Entry, error) {
	resp, err := s.srv.Spreadsheets.Values.Get(s.spreadsheetID, s.readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	var entries []Entry
	for i, row := range resp.Values {
		if i == 0 {
			// Skip the header row
			continue
		}
		if len(row) < 6 {
			continue // Skip rows that do not have all columns
		}
		entry := Entry{
			Timestamp:    fmt.Sprintf("%v", row[0]),
			EmailAddress: fmt.Sprintf("%v", row[1]),
			Winner:       fmt.Sprintf("%v", row[2]),
			SecondPlace:  fmt.Sprintf("%v", row[3]),
			ThirdPlace:   fmt.Sprintf("%v", row[4]),
			FourthPlace:  fmt.Sprintf("%v", row[5]),
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
