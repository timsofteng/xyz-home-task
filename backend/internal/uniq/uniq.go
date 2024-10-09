package uniq

import (
	"time"

	"github.com/google/uuid"
)

func New() int64 {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	id := currentTimestamp + int64(uniqueID)

	return id
}
