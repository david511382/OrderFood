package menu

import (
	"github.com/jmoiron/sqlx"
)

type MenuDb struct {
	Connect func() (*sqlx.DB, error)
}
