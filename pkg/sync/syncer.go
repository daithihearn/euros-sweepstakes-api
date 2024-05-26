package sync

import (
	"euros-sweepstakes-api/pkg/score"
	"euros-sweepstakes-api/pkg/sheets"
	"log"
)

type Syncer struct {
	ScoreService *score.Service
	SheetService *sheets.SheetService
}

func (s *Syncer) Sync() error {
	log.Println("Starting to sync...")

	// Get the entries from the sheet
	entries, err := s.SheetService.GetEntries()
	if err != nil {
		return err
	}

	// TODO: Need to call to an Odds service

	// Parse the entries into scores
	scores := make([]score.Score, 0)
	for _, entry := range entries {
		s, err := sheets.ParseScoreFromEntry(entry)
		if err != nil {
			log.Printf("Failed to parse score from entry: %v", err)
			continue
		}
		scores = append(scores, s)
	}

	// Update the cache with the new scores
	err = s.ScoreService.SetScores(scores)
	if err != nil {
		return err
	}

	return nil
}
