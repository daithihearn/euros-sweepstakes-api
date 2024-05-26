package score

import "euros-sweepstakes-api/pkg/cache"

type ServiceI interface {
	GetScores() ([]Score, error)
	SetScores(scores []Score) error
}

type Service struct {
	Cache cache.Cache[[]Score]
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

func (s *Service) SetScores(scores []Score) error {
	return s.Cache.Set("scores", scores, 0)
}
