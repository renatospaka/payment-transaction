package entity

import "time"

type TrailDate struct {
	createdAt	time.Time
	updatedAt	time.Time
	deletedAt	time.Time
}

func NewTrailDate() *TrailDate {
	return &TrailDate{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: time.Time{},
	}
}

func (t *TrailDate) SetCreationToToday() {
	t.createdAt = time.Now()
	t.updatedAt = time.Now()
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetCreationToDate(date time.Time) {
	t.createdAt = date
	t.updatedAt = date
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetAlterationdToToday() {
	t.updatedAt = time.Now()
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetAlterationdToDate(date time.Time) {
	t.updatedAt = date
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetDeletionToToday() {
	t.deletedAt = time.Now()
}

func (t *TrailDate) SetDeletionToDate(date time.Time) {
	t.deletedAt = date
}

func (t *TrailDate) Created() time.Time {
	return t.createdAt
}

func (t *TrailDate) Deleted() time.Time {
	return t.deletedAt
}

func (t *TrailDate) Updated() time.Time {
	return t.updatedAt
}