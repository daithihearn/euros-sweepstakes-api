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
			name: "Max points - all correct order",

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
			want: 26,
		},
		{
			name: "Max points - 3rd and 4th wrong",

			proposed: result.Result{
				Winner:      team.Team{Country: "Italy"},
				RunnerUp:    team.Team{Country: "England"},
				ThirdPlace:  team.Team{Country: "France"},
				FourthPlace: team.Team{Country: "Germany"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "Italy"},
				RunnerUp:    team.Team{Country: "England"},
				ThirdPlace:  team.Team{Country: "Germany"},
				FourthPlace: team.Team{Country: "France"},
			},
			want: 26,
		},
		{
			name: "3rd and 4th correct",

			proposed: result.Result{
				Winner:      team.Team{Country: "Portugal"},
				RunnerUp:    team.Team{Country: "Germany"},
				ThirdPlace:  team.Team{Country: "Italy"},
				FourthPlace: team.Team{Country: "Spain"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "England"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "Italy"},
				FourthPlace: team.Team{Country: "Spain"},
			},
			want: 10,
		},
		{
			name: "3rd and 4th wrong position",

			proposed: result.Result{
				Winner:      team.Team{Country: "Portugal"},
				RunnerUp:    team.Team{Country: "Germany"},
				ThirdPlace:  team.Team{Country: "Italy"},
				FourthPlace: team.Team{Country: "Spain"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "England"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "Spain"},
				FourthPlace: team.Team{Country: "Italy"},
			},
			want: 10,
		},
		{
			name: "First place correct",

			proposed: result.Result{
				Winner:      team.Team{Country: "Portugal"},
				RunnerUp:    team.Team{Country: "Italy"},
				ThirdPlace:  team.Team{Country: "Germany"},
				FourthPlace: team.Team{Country: "England"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "Portugal"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "Spain"},
				FourthPlace: team.Team{Country: "Netherlands"},
			},
			want: 8,
		},
		{
			name: "Second place correct",

			proposed: result.Result{
				Winner:      team.Team{Country: "Italy"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "Germany"},
				FourthPlace: team.Team{Country: "England"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "Portugal"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "Spain"},
				FourthPlace: team.Team{Country: "Netherlands"},
			},
			want: 8,
		},
		{
			name: "Third place correct",

			proposed: result.Result{
				Winner:      team.Team{Country: "Italy"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "Germany"},
				FourthPlace: team.Team{Country: "England"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "Portugal"},
				RunnerUp:    team.Team{Country: "Spain"},
				ThirdPlace:  team.Team{Country: "Germany"},
				FourthPlace: team.Team{Country: "Netherlands"},
			},
			want: 5,
		},
		{
			name: "Fourth place correct",

			proposed: result.Result{
				Winner:      team.Team{Country: "Italy"},
				RunnerUp:    team.Team{Country: "France"},
				ThirdPlace:  team.Team{Country: "England"},
				FourthPlace: team.Team{Country: "Germany"},
			},
			actual: result.Result{
				Winner:      team.Team{Country: "Portugal"},
				RunnerUp:    team.Team{Country: "Spain"},
				ThirdPlace:  team.Team{Country: "Netherlands"},
				FourthPlace: team.Team{Country: "Germany"},
			},
			want: 5,
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
