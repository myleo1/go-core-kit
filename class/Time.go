package class

import (
	"database/sql"
	"mizuki/project/core-kit/library/jsonkit"
	"mizuki/project/core-kit/library/timekit"
	"time"
)

// 同时继承scan和value方法
type Time struct {
	sql.NullTime
}

func (th Time) MarshalJSON() ([]byte, error) {
	if th.Valid {
		return jsonkit.JSON().Marshal(th.Time.Format(timekit.TimeLayoutAll))
	}
	return jsonkit.JSON().Marshal(nil)
}
func (th *Time) UnmarshalJSON(data []byte) error {
	var s *time.Time
	if err := jsonkit.JSON().Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		th.Valid = true
		th.Time = *s
	} else {
		th.Valid = false
	}
	return nil
}
