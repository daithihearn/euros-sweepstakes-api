package sheets

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
)

type SheetService struct {
	Srv           *sheets.Service
	SpreadsheetID string
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
		Srv:           srv,
		SpreadsheetID: spreadsheetID,
	}, nil
}
