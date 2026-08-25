package notify

import "mathrush/internal/domain"

func EventKinds(events []domain.Event) []string {
	out := []string{}
	for _, e := range events {
		out = append(out, e.Kind)
	}
	return out
}
func HasKind(events []domain.Event, k string) bool {
	for _, e := range events {
		if e.Kind == k {
			return true
		}
	}
	return false
}
func Last(events []domain.Event) (domain.Event, bool) {
	if len(events) == 0 {
		return domain.Event{}, false
	}
	return events[len(events)-1], true
}
