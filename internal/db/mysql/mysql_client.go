package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {
	db, err := sql.Open("mysql", source)

	if err != nil {
		panic(err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS parada (
            id  bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
            diametro float NOT NULL,
            latitud float NOT NULL,
            longitud float NOT NULL
    )`)

	db.Exec(`CREATE TABLE IF NOT EXISTS monopatin (
        id bigint not null auto_increment PRIMARY KEY,
        kilometros float NOT NULL,
        latitud float NOT NULL,
        longitud float NOT NULL,
        ultimo_mantenimiento  date NOT NULL,
        estado char(1) NOT NULL,
        id_parada bigint,
        FOREIGN KEY (id_parada) REFERENCES parada(id)
    )`)
	return &MySqlClient{db}
}
