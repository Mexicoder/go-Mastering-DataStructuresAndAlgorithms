package main

//exercise_first-recurring-char
//exercise_first_recurring_char

func firstRecurringChar(arr []int) int {
	seen := make(map[int]struct{})

	for _, v := range arr {
		if _, exists := seen[v]; exists {
			return v
		}
		seen[v] = struct{}{}
	}

	return -1
}
