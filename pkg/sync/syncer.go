package sync

import (
	"euros-sweepstakes-api/pkg/result"
	"euros-sweepstakes-api/pkg/score"
	"log"
)

type Syncer struct {
	ScoreService  *score.Service
	ResultService *result.Service
}

func (s *Syncer) Sync() error {
	log.Println("Starting to sync...")

	// Refresh the scores
	err := s.ScoreService.RefreshScores()
	if err != nil {
		return err
	}
	// Refresh the scores
	err = s.ResultService.RefreshResults()
	if err != nil {
		return err
	}

	return nil
}
