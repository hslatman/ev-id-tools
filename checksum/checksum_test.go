package checksum

import (
	"testing"
)

const (
	// Example from "Check Digit Calculation for Contract-IDs":
	// http://www.ochp.eu/wp-content/uploads/2014/02/E-Mobility-IDs_EVCOID_Check-Digit-Calculation_Explanation.pdf
	validContractID1 = "DE83DUIEN83QGZD"

	// Example from Electric Vehicle ICT Interface Specifications V1.0, Part 2: Business Objects; section 1.2.2.1:
	// http://emi3group.com/wp-content/uploads/sites/5/2018/12/eMI3-standard-v1.0-Part-2.pdf
	validContractID2 = "DE-8AA-CA2B3C4D5-L"
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
			contractID: validContractID1,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "ok/dashes",
			contractID: validContractID2,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "ok/lowercase",
			contractID: "de-8aa-ca2b3c4d5-l",
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
			contractID: validContractID1,
			want:       "D",
			wantErr:    false,
		},
		{
			name:       "ok/dashes-without-check-digit",
			contractID: "DE-8AA-CA2B3C4D5",
			want:       "L",
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
