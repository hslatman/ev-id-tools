package checksum

import (
	"fmt"
	"strings"
)

const (
	lengthExcludingCheckDigit = 14
	lengthIncludingCheckDigit = 15
)

// Verify verifies the contractID has the correct check digit
func Verify(contractID string) (bool, error) {
	// 1. perform contract format validation (based on regex?)
	// 2. perform some other sanity checks?

	normalizedContractID := normalize(contractID)

	if len(normalizedContractID) != lengthIncludingCheckDigit {
		return false, fmt.Errorf("normalized contract ID does not include check digit")
	}

	checkDigit, err := CalculateCheckDigit(normalizedContractID)
	if err != nil {
		return false, fmt.Errorf("could not calculate check digit: %w", err)
	}

	return string(normalizedContractID[14]) == checkDigit, nil
}

// CalculateCheckDigit calculates the check digit for a contract ID
// It supports contract IDs in- and excluding the check digit. It simply
// ignores the last character in the first case.
func CalculateCheckDigit(contractID string) (string, error) {

	normalizedContractID := normalize(contractID)

	if len(normalizedContractID) < lengthExcludingCheckDigit || len(normalizedContractID) > lengthIncludingCheckDigit {
		return "", fmt.Errorf("normalized contract ID %s has invalid length: %d", normalizedContractID, len(normalizedContractID))
	}

	return calculateCheckDigit(normalizedContractID)
}

// normalize normalizes contract IDs by removing dashes and
// converting lowercase characters to uppercase.
func normalize(contractID string) string {
	return strings.ToUpper(strings.ReplaceAll(contractID, "-", ""))
}

// calculateCheckDigit calculates the Contract ID Check Digit according
// to the algorithm as described in "Check Digit Calculation for Contract-IDs"
// It is available here: http://www.ochp.eu/wp-content/uploads/2014/02/E-Mobility-IDs_EVCOID_Check-Digit-Calculation_Explanation.pdf
// This implementation is based on the one in http://www.ochp.eu/id-validator/.
func calculateCheckDigit(contractID string) (string, error) {

	index := 0
	var m [4 * lengthExcludingCheckDigit]int
	for i := 0; i < lengthExcludingCheckDigit; i++ {
		for j := 0; j < 4; j++ {
			m[index] = alpha[string(contractID[i])][j]
			index++
		}
	}

	var c1, c2, c3, c4 int
	for i := 0; i < lengthExcludingCheckDigit; i++ {
		c1 += m[i*4]*p1[i][0] + m[i*4+1]*p1[i][2]
		c2 += m[i*4]*p1[i][1] + m[i*4+1]*p1[i][3]
		c3 += m[i*4+2]*p2[i][0] + m[i*4+3]*p2[i][2]
		c4 += m[i*4+2]*p2[i][1] + m[i*4+3]*p2[i][3]
	}

	q1, q2 := c1%2, c2%2

	r1 := calculateR1(c4)
	r2 := calculateR2(c3, r1)

	v := q1 + q2*2 + r1*4 + r2*16

	digit := reverse[v]

	return digit, nil
}

func calculateR1(c4 int) int {
	v := c4 % 3
	switch v {
	case 0:
		return 0
	case 1:
		return 2
	case 2:
		return 1
	default:
		return -1 // satisfy return; never happens
	}
}

func calculateR2(c3, r1 int) int {
	v := (c3 % 3) + r1
	switch v {
	case 0:
		return 0
	case 1:
		return 2
	case 2:
		return 1
	case 3:
		return 0
	case 4:
		return 2
	default:
		return -1 // satisfy return; never happens
	}
}
