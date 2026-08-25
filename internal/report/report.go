package report

import (
	"mathrush/internal/domain"
	"sort"
)

func Total(records []domain.Record) int {
	n := 0
	for _, r := range records {
		n += r.Score
	}
	return n
}
func ByScore(records []domain.Record) []domain.Record {
	out := append([]domain.Record(nil), records...)
	sort.Slice(out, func(i, j int) bool { return out[i].Score > out[j].Score })
	return out
}
func Summary(u domain.User) string { return u.Name + ":" + domain.Rank(u.Score) }
func StatusCounts(records []domain.Record) map[string]int {
	m := map[string]int{}
	for _, r := range records {
		m[r.Status]++
	}
	return m
}
