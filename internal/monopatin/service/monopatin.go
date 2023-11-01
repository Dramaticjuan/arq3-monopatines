package service

import (
	repo "github.com/Dramaticjuan/arq3-monopatines/internal/monopatin/repo/mysql"
)

type MonopatinService struct {
	repo *repo.MonopatinRepo
}

func NewMonopatinService(repo *repo.MonopatinRepo) *MonopatinService {
	return &MonopatinService{
		repo: repo,
	}
}

func (ms *MonopatinService) UpdateParada(id uint, id_parada uint) error {
	return ms.repo.UpdateParada(id, id_parada)
}
