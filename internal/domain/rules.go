package domain

func IsPerfect(r Record) bool { return r.Answer == 5 && r.Score == 10 }
func NeedsReview(r Record) bool {
	if r.Score < 5 {
		return true
	}
	return r.Status == "pending"
}
func Rank(score int) string {
	if score >= 100 {
		return "gold"
	}
	if score >= 50 {
		return "silver"
	}
	return "bronze"
}
func MergeUser(u User, r Record) User {
	u.Attempts++
	u.Score += r.Score
	u.UpdatedAt = r.UpdatedAt
	return u
}
func EventFor(r Record, kind string) Event {
	return Event{ID: r.ID + "-" + kind, RecordID: r.ID, Kind: kind, Payload: r.Status, At: r.UpdatedAt}
}
func AuditFor(r Record, actor, action string) Audit {
	return Audit{ID: r.ID + "-" + action, RecordID: r.ID, Actor: actor, Action: action, Detail: r.Status, At: r.UpdatedAt}
}
