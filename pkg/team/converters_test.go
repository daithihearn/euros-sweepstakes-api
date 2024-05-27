package team

import (
	"testing"
)

func TestParseTeamFromString(t *testing.T) {
	tests := []struct {
		name    string
		country string
		odds    string
		want    Team
		wantErr bool
	}{
		{
			name:    "valid team and flag",
			country: "Italy",
			odds:    "10/1",
			want:    Team{Country: "Italy", Odds: "10/1"},
			wantErr: false,
		},
		{
			name:    "empty team",
			country: "",
			odds:    "",
			want:    Team{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTeam(tt.country, tt.odds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
