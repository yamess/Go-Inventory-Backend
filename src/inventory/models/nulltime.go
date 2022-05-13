package models

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
)

type NullTime sql.NullTime

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{Time: t.Time.UTC(), Valid: false}
	} else {
		*nt = NullTime{Time: t.Time.UTC(), Valid: true}
	}

	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
