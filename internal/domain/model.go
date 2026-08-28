package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Record struct {
	ID, UserID, Expression, Status string
	Answer, Score, Version         int
	CreatedAt, UpdatedAt           time.Time
}
type User struct {
	ID, Name        string
	Score, Attempts int
	UpdatedAt       time.Time
}
type Event struct {
	ID, RecordID, Kind, Payload string
	At                          time.Time
}
type Audit struct {
	ID, RecordID, Actor, Action, Detail string
	At                                  time.Time
}

func ValidateRecord(r Record) error {
	if strings.TrimSpace(r.ID) == "" {
		return errors.New("id required")
	}
	if r.UserID == "" {
		return errors.New("user required")
	}
	if r.Expression == "" {
		return errors.New("expression required")
	}
	if r.Answer < 0 || r.Answer > 5 {
		return errors.New("answer out of range")
	}
	return nil
}
func ScoreAnswer(answer int) int {
	if answer == 5 {
		return 10
	}
	if answer >= 3 {
		return 5
	}
	return 0
}
func NewRecord(id, user, expr string, answer int) Record {
	now := time.Now().UTC()
	return Record{ID: id, UserID: user, Expression: expr, Answer: answer, Score: ScoreAnswer(answer), Status: "pending", Version: 1, CreatedAt: now, UpdatedAt: now}
}
func (r *Record) Approve() error {
	if r.Status != "pending" {
		return fmt.Errorf("cannot approve %s", r.Status)
	}
	r.Status = "approved"
	r.Version++
	r.UpdatedAt = time.Now().UTC()
	return nil
}
func (r *Record) Archive() error {
	if r.Status != "approved" {
		return fmt.Errorf("cannot archive %s", r.Status)
	}
	r.Status = "archived"
	r.Version++
	r.UpdatedAt = time.Now().UTC()
	return nil
}
func (r *Record) Confirm() error {
	if r.Status == "archived" {
		return errors.New("archived")
	}
	r.Status = "confirmed"
	r.Version++
	r.UpdatedAt = time.Now().UTC()
	return nil
}
func NormalizeName(s string) string { return strings.TrimSpace(strings.ToUpper(s)) }
func ValidStatus(s string) bool {
	switch s {
	case "pending", "approved", "archived", "confirmed":
		return true
	}
	return false
}
