package calculate

import (
	"fmt"
)

func CalculateNotes(amount int) (map[int]int, error) {
	notes := []int{200, 100, 50, 20, 10, 5}
	noteList := make(map[int]int)
	remainingAmount := amount

	for _, note := range notes {
		if remainingAmount >= note {
			noteList[note] = remainingAmount / note
			remainingAmount %= note
		}
	}

	if remainingAmount != 0 {
		return nil, fmt.Errorf("Valor indisponível para saque.")
	}

	return noteList, nil
}
