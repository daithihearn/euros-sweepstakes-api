package sheets

import (
	"euros-sweepstakes-api/pkg/score"
	"testing"
)

func TestParseNameFromEmail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid email",
			input:   "john.smith@mail.com",
			want:    "John Smith",
			wantErr: false,
		},
		{
			name:    "invalid email",
			input:   "notanemail",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNameFromEmail(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNameFromEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNameFromEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTeamFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    score.Team
		wantErr bool
	}{
		{
			name:    "valid team and flag",
			input:   "🇮🇹 Italy",
			want:    score.Team{Flag: "🇮🇹", Country: "Italy"},
			wantErr: false,
		},
		{
			name:    "invalid team no flag",
			input:   "Italy",
			want:    score.Team{},
			wantErr: true,
		},
		{
			name:    "empty team",
			input:   "",
			want:    score.Team{},
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
