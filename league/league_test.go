package league

import "testing"

func TestLeague_populateMatches(t *testing.T) {
	type fields struct {
		players Players
		matches []Match
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &League{
				players: tt.fields.players,
				matches: tt.fields.matches,
			}
			if err := l.populateMatches(); (err != nil) != tt.wantErr {
				t.Errorf("League.populateMatches() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
