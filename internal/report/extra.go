package report

import "mathrush/internal/domain"

func Average(records []domain.Record) float64 {
	if len(records) == 0 {
		return 0
	}
	return float64(Total(records)) / float64(len(records))
}
func PerfectCount(records []domain.Record) int {
	n := 0
	for _, r := range records {
		if domain.IsPerfect(r) {
			n++
		}
	}
	return n
}
func Pending(records []domain.Record) []domain.Record {
	out := []domain.Record{}
	for _, r := range records {
		if r.Status == "pending" {
			out = append(out, r)
		}
	}
	return out
}
func UserScore(u domain.User) int { return u.Score }
