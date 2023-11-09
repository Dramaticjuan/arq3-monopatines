package model

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Monopatin struct {
	ID                  int
	Kilometros          float64
	Latitud             float64
	Longitud            float64
	UltimoMantenimiento time.Time
	Estado              string
    IDParada            null.Int
}
