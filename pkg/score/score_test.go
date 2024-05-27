package score

import (
	"euros-sweepstakes-api/pkg/result"
	"euros-sweepstakes-api/pkg/team"
	"testing"
)

func TestCalculateTotalScore(t *testing.T) {
	tests := []struct {
		name     string
		proposed result.Result
		actual   result.Result
		want     int
	}{
		{
			name: "Max points",

			proposed: result.Result{
				Winner:      team.Team{Country: "Italy"},
				RunnerUp:    team.Team{Country: "England"},
				ThirdPlace:  team.Team{Country: "France"},
				FourthPlace: team.Team{Country: "Germany"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "Italy"},
				RunnerUp:    team.Team{Country: "England"},
				ThirdPlace:  team.Team{Country: "France"},
				FourthPlace: team.Team{Country: "Germany"},
			},
			want: 32,
		},
		{
			name: "Some points",

			proposed: result.Result{
				Winner:      team.Team{Country: "England"},
				RunnerUp:    team.Team{Country: "Germany"},
				ThirdPlace:  team.Team{Country: "France"},
				FourthPlace: team.Team{Country: "Italy"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "England"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "Spain"},
				FourthPlace: team.Team{Country: "Italy"},
			},
			want: 21,
		},
		{
			name: "No points",
			proposed: result.Result{
				Winner:      team.Team{Country: "England"},
				RunnerUp:    team.Team{Country: "Germany"},
				ThirdPlace:  team.Team{Country: "France"},
				FourthPlace: team.Team{Country: "Italy"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "Spain"},
				RunnerUp:    team.Team{Country: "Portugal"},
				ThirdPlace:  team.Team{Country: "Belgium"},
				FourthPlace: team.Team{Country: "Netherlands"},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateTotalScore(tt.proposed, tt.actual)
			if got != tt.want {
				t.Errorf("CalculateTotalScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
