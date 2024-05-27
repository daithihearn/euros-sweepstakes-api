package team

import (
	"testing"
)

func TestParseTeamFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Team
		wantErr bool
	}{
		{
			name:    "valid team and flag",
			input:   "Italy",
			want:    Team{Country: "Italy"},
			wantErr: false,
		},
		{
			name:    "empty team",
			input:   "",
			want:    Team{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTeamFromString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTeamFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseTeamFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
