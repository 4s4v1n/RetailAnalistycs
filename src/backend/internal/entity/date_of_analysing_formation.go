package entity

import (
	"APG6/internal/entity/utils"
	"encoding/json"
	"time"
)

type DateOfAnalysingFormation struct {
	Date time.Time `json:"date,time.RFC3339" csv:"date" db:"date"`
}

func (df *DateOfAnalysingFormation) UnmarshalJSON(b []byte) error {
	type formationAlias DateOfAnalysingFormation
	alias := &struct {
		*formationAlias
		Date string `json:"date"`
	}{
		formationAlias: (*formationAlias)(df),
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	t, err := utils.ParseDatetime(alias.Date)
	if err != nil {
		return err
	}
	df.Date = t
	return nil
}

type UpdateFormation struct {
	OldDate time.Time `json:"old_date,time.RFC3339" csv:"-" db:"-"`
	NewDate time.Time `json:"new_date,time.RFC3339" csv:"-" db:"-"`
}

func (uf *UpdateFormation) UnmarshalJSON(b []byte) error {
	type updateAlias UpdateFormation
	alias := &struct {
		*updateAlias
		OldDate string `json:"old_date"`
		NewDate string `json:"new_date"`
	}{
		updateAlias: (*updateAlias)(uf),
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	oldDate, err := utils.ParseDatetime(alias.OldDate)
	if err != nil {
		return err
	}
	newDate, err := utils.ParseDatetime(alias.NewDate)
	if err != nil {
		return err
	}

	uf.OldDate = oldDate
	uf.NewDate = newDate

	return nil
}
