package main

import (
	"fmt"
	"net/http"
	"os"

	database "github.com/Dramaticjuan/arq3-monopatines/internal/db/mysql"
	repo "github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/repo/mysql"
	"github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/rest"
	"github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	fmt.Println(conn)

	db := database.NewSqlClient(conn)
	if db == nil {
		panic("algo falló")
	}

	repom := repo.NewMonopatinRepo(db)
	repop := repo.NewParadaRepo(db)

	ms := service.NewMonopatinService(repom, repop)

	restm := rest.NewMonopatinController(repom, ms)
	restp := rest.NewParadaController(repop)

	routes := Routes(*restm, *restp)

	http.ListenAndServe(":3000", routes)
}

func Routes(m rest.MonopatinController, p rest.ParadaController) *chi.Mux {
	router := chi.NewMux()
	router.Route("/monopatin", func(r chi.Router) {
		r.Get("/{id}", m.GetMonopatin)
		r.Get("/", m.GetAllMonopatin)
		r.Post("/", m.CreateMonopatin)
		r.Delete("/{id}", m.DeleteMonopatin)
		r.Get("/{latitud}/{longitud}/{rango}", m.GetAllMonopatinCercanos)
		r.Patch("/{id}/{latitud}/{longitud}/{kilometros}", m.UpdateKilometrosYCoordenadas)
		r.Patch("/{id}/{estado}", m.UpdateEstado)
		r.Patch("/{id}/parada/{parada}", m.UpdateParada)
	})
	router.Route("/parada", func(r chi.Router) {
		r.Post("/", p.CreateParada)
		r.Delete("/{id}", p.DeleteParada)
		r.Get("/", p.GetAllParada)
		r.Get("/{id}", p.GetParada)
	})
	return router
}
