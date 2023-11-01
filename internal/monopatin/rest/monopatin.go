package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dramaticjuan/arq3-monopatines/internal/model"
	repo "github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/repo/mysql"
	"github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type MonopatinController struct {
	repo *repo.MonopatinRepo
	ms   *service.MonopatinService
}

func NewMonopatinController(repo *repo.MonopatinRepo, ms *service.MonopatinService) *MonopatinController {
	return &MonopatinController{
		repo: repo,
		ms:   ms,
	}
}

func (mc *MonopatinController) CreateMonopatin(w http.ResponseWriter, r *http.Request) {
	var monopatinReq model.Monopatin
	err := json.NewDecoder(r.Body).Decode(&monopatinReq)
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), 400)
		return
	}
	err = mc.repo.CreateMonopatin(monopatinReq)
	if err != nil {
		http.Error(w, "Internal error"+err.Error(), 500)
		return
	}
	render.PlainText(w, r, "Monopatín creado.")
}

func (mc *MonopatinController) DeleteMonopatin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	number, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	mc.repo.DeleteMonopatin(uint(number))
}

func (mc *MonopatinController) GetMonopatin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	number, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	monopatin, er := mc.repo.GetMonopatin(uint(number))
	if er != nil {
		http.Error(w, "Internal error: "+er.Error(), 500)
		return
	}
	render.JSON(w, r, monopatin)
}

func (mc *MonopatinController) GetAllMonopatin(w http.ResponseWriter, r *http.Request) {
	all, err := mc.repo.ListMonopatines()
	if err != nil {
		http.Error(w, "Internal error"+err.Error(), 500)
		return
	}
	render.JSON(w, r, all)
}

func (mc *MonopatinController) GetAllMonopatinCercanos(w http.ResponseWriter, r *http.Request) {
	latitud_param := chi.URLParam(r, "latitud")
	latitud, err1 := strconv.ParseFloat(latitud_param, 64)
	if err1 != nil {
		http.Error(w, "Bad Request "+err1.Error(), 400)
		return
	}
	longitud_param := chi.URLParam(r, "longitud")
	longitud, err2 := strconv.ParseFloat(longitud_param, 64)
	if err2 != nil {
		http.Error(w, "Bad Request "+err2.Error(), 400)
		return
	}
	rango_param := chi.URLParam(r, "rango")
	rango, err3 := strconv.ParseFloat(rango_param, 64)
	if err3 != nil {
		http.Error(w, "Bad Request "+err3.Error(), 400)
		return
	}

	all, err := mc.repo.ListMonopatinesCercanos(latitud, longitud, rango)
	if err != nil {
		http.Error(w, "Internal error"+err.Error(), 500)
		return
	}
	render.JSON(w, r, all)
}

func (mc *MonopatinController) GetUltimoMonopatin(w http.ResponseWriter, r *http.Request) {
	monopatin, err := mc.repo.UltimoAgregado()
	if err != nil {
		http.Error(w, "Internal error: "+err.Error(), 500)
		return
	}
	render.JSON(w, r, monopatin)
}

func (mc *MonopatinController) UpdateKilometrosYCoordenadas(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	number, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	latitud_param := chi.URLParam(r, "latitud")
	latitud, err1 := strconv.ParseFloat(latitud_param, 64)
	if err1 != nil {
		http.Error(w, "Bad Request "+err1.Error(), 400)
		return
	}
	longitud_param := chi.URLParam(r, "longitud")
	longitud, err2 := strconv.ParseFloat(longitud_param, 64)
	if err2 != nil {
		http.Error(w, "Bad Request "+err2.Error(), 400)
		return
	}
	kilometros_param := chi.URLParam(r, "rango")
	kilometros, err3 := strconv.ParseFloat(kilometros_param, 64)
	if err3 != nil {
		http.Error(w, "Bad Request "+err3.Error(), 400)
		return
	}

	errU := mc.repo.UpdateKilometrosYCoordenadas(uint(number), kilometros, latitud, longitud)
	if errU != nil {
		http.Error(w, "Internal error: "+errU.Error(), 500)
		return
	}
	monopatin, errQ := mc.repo.GetMonopatin(uint(number))
	if errQ != nil {
		http.Error(w, "Internal error: "+errQ.Error(), 500)
		return
	}
	render.JSON(w, r, monopatin)
}

func (mc *MonopatinController) UpdateParada(w http.ResponseWriter, r *http.Request) {
	id_param := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(id_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	id_parada_param := chi.URLParam(r, "latitud")
	id_parada, err1 := strconv.ParseFloat(id_parada_param, 64)
	if err1 != nil {
		http.Error(w, "Bad Request "+err1.Error(), 400)
		return
	}

	err = mc.ms.UpdateParada(uint(id), uint(id_parada))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	monopatin, errQ := mc.repo.GetMonopatin(uint(id))
	if errQ != nil {
		http.Error(w, "Internal error: "+errQ.Error(), 500)
		return
	}
	render.JSON(w, r, monopatin)
}

func (mc *MonopatinController) UpdateEstado(w http.ResponseWriter, r *http.Request) {
	// TODO: chequear rol
	id_param := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(id_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	estado := chi.URLParam(r, "estado")

	err = mc.repo.UpdateEstado(uint(id), estado)
	if err != nil {
		http.Error(w, "Internal error: "+err.Error(), 500)
		return
	}
	monopatin, errQ := mc.repo.GetMonopatin(uint(id))
	if errQ != nil {
		http.Error(w, "Internal error: "+errQ.Error(), 500)
		return
	}
	render.JSON(w, r, monopatin)
}
