package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dramaticjuan/arq3-monopatines/internal/model"
	repo "github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/repo/mysql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ParadaController struct {
	repo *repo.ParadaRepo
}

func NewParadaController(repo *repo.ParadaRepo) *ParadaController {
	return &ParadaController{
		repo: repo,
	}
}

func (pc ParadaController) CreateParada(w http.ResponseWriter, r *http.Request) {
	var paradaReq model.Parada
	err := json.NewDecoder(r.Body).Decode(&paradaReq)
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), 400)
		return
	}
	err = pc.repo.CreateParada(paradaReq)
	if err != nil {
		http.Error(w, "Internal error"+err.Error(), 500)
		return
	}
	render.PlainText(w, r, "Parada creada.")
}

func (pc ParadaController) DeleteParada(w http.ResponseWriter, r *http.Request) {
	id_param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	pc.repo.DeleteParada(id)
}

func (pc ParadaController) GetParada(w http.ResponseWriter, r *http.Request) {
	id_param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	p, er := pc.repo.GetParada(id)
	if er != nil {
		http.Error(w, "Internal error "+er.Error(), 500)
	}
	render.JSON(w, r, p)
}

func (pc ParadaController) GetAllParada(w http.ResponseWriter, r *http.Request) {
	all, err := pc.repo.ListParada()
	if err != nil {
		http.Error(w, "Internal error "+err.Error(), 500)
		return
	}
	render.JSON(w, r, all)
}
