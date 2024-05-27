package result

import (
	"euros-sweepstakes-api/pkg/team"
)

type Result struct {
	Winner      team.Team `json:"winner"`
	RunnerUp    team.Team `json:"runnerUp"`
	ThirdPlace  team.Team `json:"thirdPlace"`
	FourthPlace team.Team `json:"fourthPlace"`
}
