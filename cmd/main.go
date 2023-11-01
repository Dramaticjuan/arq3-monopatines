package main

import (
	"fmt"
	"os"

	database "github.com/Dramaticjuan/arq3-monopatines/internal/db/mysql"
	repo "github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/repo/mysql"
	"github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/rest"
	"github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/service"
)

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	fmt.Printf(conn)

	db := database.NewSqlClient(conn)
	if db == nil {
		panic("algo falló")
	}

	repo := repo.NewMonopatinRepo(db)

	service := service.NewMonopatinService(repo)

	rest := rest.NewMonopatinController(repo, service)

}
