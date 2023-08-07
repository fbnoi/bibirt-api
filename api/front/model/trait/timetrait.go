package trait

import (
	"time"
)

type TimeTrait struct {
	CreatedAt int64
	UpdatedAt int64
}

func (tt *TimeTrait) GetCreatedAt() time.Time {
	return time.Unix(tt.CreatedAt, 0)
}

func (tt *TimeTrait) GetUpdatedAt() time.Time {
	return time.Unix(tt.UpdatedAt, 0)
}
