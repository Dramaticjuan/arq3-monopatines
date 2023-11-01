package model

import (
	"database/sql"
	"time"
)

type Monopatin struct {
	ID                  uint64
	Kilometros          float64
	Latitud             float64
	Longitud            float64
	UltimoMantenimiento time.Time
	Estado              string
	IDParada            sql.NullInt32
}
