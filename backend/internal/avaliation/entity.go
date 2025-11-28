package avaliation

import (
	"encoding/json"
	"time"
)

type avaliation struct {
	ID           string          `db:"id" json:"id"`
	LastUpdated  *time.Time      `db:"last_updated" json:"last_updated,omitempty"`
	Active       bool            `db:"active" json:"active"`
	CreatedAt    time.Time       `db:"created_at" json:"created_at"`
	ResourceJSON json.RawMessage `db:"resource_json" json:"resource_json"`
}
