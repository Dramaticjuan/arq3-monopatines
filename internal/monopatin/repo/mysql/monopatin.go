package repo

import (
	database "github.com/Dramaticjuan/arq3-monopatines/internal/db/mysql"
	"github.com/Dramaticjuan/arq3-monopatines/internal/model"
)

type MonopatinRepo struct {
	db *database.MySqlClient
}

func NewMonopatinRepo(db *database.MySqlClient) *MonopatinRepo {
	return &MonopatinRepo{
		db: db,
	}
}

const createMonopatin = `-- name: CreateMonopatin :exec
INSERT INTO monopatin (
  latitud, longitud, ultimo_mantenimiento, kilometros, estado
) VALUES (?,?,?,?,?)
`

func (mr *MonopatinRepo) CreateMonopatin(m model.Monopatin) error {
	stmt, err := mr.db.Prepare(createMonopatin)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.Latitud, m.Longitud, m.UltimoMantenimiento, m.Kilometros, m.Estado)
	return err
}

const deleteMonopatin = `-- name: DeleteMonopatin :exec
DELETE FROM monopatin WHERE id = ?
`

func (mr *MonopatinRepo) DeleteMonopatin(id int64) error {

	stmt, err := mr.db.Prepare(deleteMonopatin)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

const getMonopatin = `-- name: GetMonopatin :one
SELECT id, kilometros, latitud, longitud, ultimo_mantenimiento, estado, id_parada FROM monopatin WHERE id = ?
`

func (mr *MonopatinRepo) GetMonopatin(id int64) (*model.Monopatin, error) {
	i := model.Monopatin{}
	row, err := mr.db.Query(getMonopatin, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(
		&i.ID,
		&i.Kilometros,
		&i.Latitud,
		&i.Longitud,
		&i.UltimoMantenimiento,
		&i.Estado,
		&i.IDParada,
	)
	if err != nil {
		return nil, err
	}
	return &i, err
}

const listMonopatines = `-- name: ListMonopatines :many
SELECT id, kilometros, latitud, longitud, ultimo_mantenimiento, estado, id_parada FROM monopatin
`

func (mr *MonopatinRepo) ListMonopatines() ([]*model.Monopatin, error) {
	rows, err := mr.db.Query(listMonopatines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	monopatines := []*model.Monopatin{}
	for rows.Next() {
		var i model.Monopatin
		if err := rows.Scan(
			&i.ID,
			&i.Kilometros,
			&i.Latitud,
			&i.Longitud,
			&i.UltimoMantenimiento,
			&i.Estado,
			&i.IDParada,
		); err != nil {
			return nil, err
		}
		monopatines = append(monopatines, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return monopatines, nil
}

const listMonopatinesCercanos = `-- name: ListMonopatinesCercanos :many
SELECT id, kilometros, latitud, longitud, ultimo_mantenimiento, estado, id_parada
FROM monopatin m
WHERE 
       ( ( ( Acos(Sin(( ? * Pi() / 180 )) * Sin((
                  m.latitud* Pi() / 180 )) +
                    Cos
                      ((
                        ? * Pi() / 180 )) * Cos((
                    m.latitud* Pi() / 180 )) *
                    Cos
                      ((
                        (
                             ? - m.longitud ) * Pi() / 180 ))) ) *
           180 / Pi
           ()
         ) * 60 * 1.1515 * 1.609344 * 1000 ) <=?
         AND id_parada IS NOT NULL
         AND estado= 'h';
`

func (mr *MonopatinRepo) ListMonopatinesCercanos(latitud float64, longitud float64, rango float64) ([]*model.Monopatin, error) {
	rows, err := mr.db.Query(listMonopatinesCercanos, latitud, latitud, longitud, rango)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	monopatines := []*model.Monopatin{}
	for rows.Next() {
		var i model.Monopatin
		if err := rows.Scan(
			&i.ID,
			&i.Kilometros,
			&i.Latitud,
			&i.Longitud,
			&i.UltimoMantenimiento,
			&i.Estado,
			&i.IDParada,
		); err != nil {
			return nil, err
		}
		monopatines = append(monopatines, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return monopatines, nil
}

const ultimoAgregado = `-- name: UltimoAgregado :one
SELECT id, kilometros, latitud, longitud, ultimo_mantenimiento, estado, id_parada FROM monopatin WHERE id= LAST_INSERT_ID()
`

func (mr *MonopatinRepo) UltimoAgregado() (*model.Monopatin, error) {
	row:= mr.db.QueryRow(ultimoAgregado)
	i := model.Monopatin{}
    err := row.Scan(
		&i.ID,
		&i.Kilometros,
		&i.Latitud,
		&i.Longitud,
		&i.UltimoMantenimiento,
		&i.Estado,
		&i.IDParada,
	)
	return &i, err
}

const UpdateKilometrosYCoordenadas = `-- name: UpdateKilometros :exec
UPDATE monopatin SET kilometros= kilometros+?, latitud=?, longitud=? WHERE id=?
`

func (mr *MonopatinRepo) UpdateKilometrosYCoordenadas(id int64, km float64, latitud float64, longitud float64) error {
	_, err := mr.db.Exec(UpdateKilometrosYCoordenadas, km, latitud, longitud, id)
	return err
}

const updateParada = `-- name: UpdateParada :exec
UPDATE monopatin SET id_parada= ? WHERE id=?
`

func (mr *MonopatinRepo) UpdateParada(id int64, id_parada int64) error {
	_, err := mr.db.Exec(updateParada, id_parada, id)
	return err
}

const updateEstado = `-- name: UpdateEstado :exec
UPDATE monopatin SET estado= ? WHERE id=?
`

func (mr *MonopatinRepo) UpdateEstado(id int64, estado string) error {
	_, err := mr.db.Exec(updateEstado, id, estado)
	return err
}
