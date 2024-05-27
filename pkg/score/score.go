package score

import "euros-sweepstakes-api/pkg/result"

type Score struct {
	Name       string        `json:"name"`
	Result     result.Result `json:"result"`
	TotalScore int           `json:"totalScore"`
}
