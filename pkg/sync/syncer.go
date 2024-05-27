package sync

import (
	"euros-sweepstakes-api/pkg/result"
	"euros-sweepstakes-api/pkg/score"
	"log"
)

type Syncer struct {
	ScoreSync  *score.Sync
	ResultSync *result.Sync
}

func (s *Syncer) Sync() error {
	log.Println("Starting to sync...")

	// Refresh the scores
	err := s.ScoreSync.RefreshScores()
	if err != nil {
		return err
	}
	// Refresh the scores
	err = s.ResultSync.RefreshResults()
	if err != nil {
		return err
	}

	log.Println("Sync complete")
	return nil
}
