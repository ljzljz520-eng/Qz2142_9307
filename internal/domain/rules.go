package domain

import "strconv"

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

// EventFor builds an event describing a transition of r. The ID embeds r.Version
// so that successive results for the same record (e.g. multiple confirms) land on
// distinct keys instead of overwriting one another — earlier and later results
// are both preserved as separate records.
func EventFor(r Record, kind string) Event {
	return Event{ID: r.ID + "-" + kind + "-" + strconv.Itoa(r.Version), RecordID: r.ID, Kind: kind, Payload: r.Status, At: r.UpdatedAt}
}
func AuditFor(r Record, actor, action string) Audit {
	return Audit{ID: r.ID + "-" + action + "-" + strconv.Itoa(r.Version), RecordID: r.ID, Actor: actor, Action: action, Detail: r.Status, At: r.UpdatedAt}
}
