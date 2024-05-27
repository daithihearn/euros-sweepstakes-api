package result

import (
	"euros-sweepstakes-api/pkg/cache"
)

type ServiceI interface {
	GetResult() (Result, error)
	SetResult(Result) error
}

type Service struct {
	Cache cache.Cache[Result]
}

func (s *Service) GetResult() (Result, error) {
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

func (s *Service) SetResult(result Result) error {
	err := s.Cache.Set("results", result, 0)
	if err != nil {
		return err
	}
	return nil
}
