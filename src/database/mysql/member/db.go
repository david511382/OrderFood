package member

import (
	"github.com/jmoiron/sqlx"
)

type MemberDb struct {
	Connect func() (*sqlx.DB, error)
}
