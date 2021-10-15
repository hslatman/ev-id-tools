package checksum

import (
	"testing"
)

func TestVerify(t *testing.T) {
	tests := []struct {
		name       string
		contractID string
		want       bool
		wantErr    bool
	}{
		{
			name:       "ok",
			contractID: "DE83DUIEN83QGZD", // TODO: reference source
			want:       true,
			wantErr:    false,
		},
		{
			name:       "fail/too-short",
			contractID: "DE83DUIEN83QG",
			want:       false,
			wantErr:    true,
		},
		{
			name:       "fail/too-long",
			contractID: "DE83DUIEN83QGZXX",
			want:       false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.contractID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateCheckDigit(t *testing.T) {
	tests := []struct {
		name       string
		contractID string
		want       string
		wantErr    bool
	}{
		{
			name:       "ok",
			contractID: "DE83DUIEN83QGZ", // TODO: reference source
			want:       "D",
			wantErr:    false,
		},
		{
			name:       "fail/too-short",
			contractID: "DE83DUIEN83QG",
			want:       "",
			wantErr:    true,
		},
		{
			name:       "fail/too-long",
			contractID: "DE83DUIEN83QGZXX",
			want:       "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateCheckDigit(tt.contractID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateCheckDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
