package repo

import (
	database "github.com/Dramaticjuan/arq3-monopatines/internal/db/mysql"
	"github.com/Dramaticjuan/arq3-monopatines/internal/model"
)

type ParadaRepo struct {
	db *database.MySqlClient
}

func NewParadaRepo(db *database.MySqlClient) *ParadaRepo {
	return &ParadaRepo{
		db: db,
	}
}

const createParada = `INSERT INTO parada (
  diametro, latitud, longitud
) VALUES (?, ?, ?);
`

func (pr *ParadaRepo) CreateParada(p model.Parada) error {
	stmt, err := pr.db.Prepare(createParada)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Diametro, p.Latitud, p.Longitud)
	return err
}

const DeleteParada = `-- name: DeleteParada :exec
DELETE FROM parada WHERE id = ?
`

func (pr *ParadaRepo) DeleteParada(id int64) error {
	stmt, err := pr.db.Prepare(DeleteParada)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

const getParada = `-- name: GetParada :one
SELECT id, diametro, latitud, longitud FROM parada WHERE id = ?
`

func (pr *ParadaRepo) GetParada(id int64) (*model.Parada, error) {
	i := model.Parada{}
	row:= pr.db.QueryRow(getParada, id)
    err := row.Scan(
		&i.ID,
		&i.Diametro,
		&i.Latitud,
		&i.Longitud,
	)
	if err != nil {
		return nil, err
	}
	return &i, err
}

const listParada = `-- name: GetParada :many
SELECT id, diametro, latitud, longitud FROM parada 
`

func (pr *ParadaRepo) ListParada() ([]*model.Parada, error) {
	row, err := pr.db.Query(listParada)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	paradas := []*model.Parada{}
	for row.Next() {
		var i model.Parada
		if err = row.Scan(
			&i.ID,
			&i.Diametro,
			&i.Latitud,
			&i.Longitud,
		); err != nil {
			return nil, err
		}
		paradas = append(paradas, &i)
	}
	if err != nil {
		return nil, err
	}
	return paradas, nil
}

const estaEnParada = `
SELECT IF(
       ( ( ( Acos(Sin(( ? * Pi() / 180 )) * Sin((
                  p.latitud* Pi() / 180 )) +
                    Cos
                      ((
                        ? * Pi() / 180 )) * Cos((
                    p.latitud* Pi() / 180 )) *
                    Cos
                      ((
                        (
                             ? - p.longitud ) * Pi() / 180 ))) ) *
           180 / Pi()
         ) * 60 * 1.1515 * 1.609344 * 1000 ) <=p.diametro, 1, 0) AS esta
FROM   parada p
WHERE p.id=?;
`

func (pr *ParadaRepo) EstaEnParada(id int64, latitud float64, longitud float64) (bool, error) {

	var esta int =10 
	row := pr.db.QueryRow(estaEnParada, latitud, latitud, longitud, id)
    err := row.Scan(
        &esta,
	)
	if err != nil {
		return false, err
	}
    if esta== 1{
        return true, nil
    }
	return false, nil
}
