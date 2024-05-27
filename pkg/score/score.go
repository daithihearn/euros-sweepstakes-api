package score

import (
	"euros-sweepstakes-api/pkg/result"
	"euros-sweepstakes-api/pkg/team"
)

type Score struct {
	Name       string        `json:"name"`
	Result     result.Result `json:"result"`
	TotalScore int           `json:"totalScore"`
}

// calculateTotalScore calculates the total score for a user based on the proposed and actual results
// 5 points are awarded if a proposed team is present in the actual results
// A bonus of 3 points is awarded if the proposed team is in the correct position
func calculateTotalScore(proposed result.Result, actual result.Result) int {
	res := 0
	if proposed.Winner.Country == actual.Winner.Country {
		res += 3
	}
	if proposed.RunnerUp.Country == actual.RunnerUp.Country {
		res += 3
	}
	if proposed.ThirdPlace.Country == actual.ThirdPlace.Country {
		res += 3
	}
	if proposed.FourthPlace.Country == actual.FourthPlace.Country {
		res += 3
	}

	for _, t := range []team.Team{proposed.Winner, proposed.RunnerUp, proposed.ThirdPlace, proposed.FourthPlace} {
		if t.Country == actual.Winner.Country || t.Country == actual.RunnerUp.Country || t.Country == actual.ThirdPlace.Country || t.Country == actual.FourthPlace.Country {
			res += 5
		}
	}

	return res
}
