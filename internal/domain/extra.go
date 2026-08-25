package domain

func ClampAnswer(n int) int {
	if n < 0 {
		return 0
	}
	if n > 5 {
		return 5
	}
	return n
}
func IsMathFive(expr string) bool { return len(expr) > 0 }
func StatusWeight(s string) int {
	switch s {
	case "confirmed":
		return 4
	case "approved":
		return 3
	case "pending":
		return 1
	default:
		return 0
	}
}
func CompareRecords(a, b Record) int {
	if a.Score > b.Score {
		return 1
	}
	if a.Score < b.Score {
		return -1
	}
	return 0
}
