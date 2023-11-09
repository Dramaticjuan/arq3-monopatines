package service

import (
	"errors"

	repo "github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/repo/mysql"
)

type MonopatinService struct {
	repo       *repo.MonopatinRepo
	repoParada *repo.ParadaRepo
}

func NewMonopatinService(repo *repo.MonopatinRepo, repoParada *repo.ParadaRepo) *MonopatinService {
	return &MonopatinService{
		repo:       repo,
		repoParada: repoParada,
	}
}

func (ms *MonopatinService) UpdateParada(id int64, id_parada int64) error {
	monopatin, err := ms.repo.GetMonopatin(id)
	if err != nil {
		return err
	}
	esta, err2 := ms.repoParada.EstaEnParada(id_parada, monopatin.Latitud, monopatin.Longitud)
	if err2 != nil {
		return err2
	}
	if esta {
		return ms.repo.UpdateParada(id, id_parada)
	}
	return errors.New("el monopatín no está en la parada")
}
