package service

import (
	"sync"

	"github.com/gouravk3/ntt-golang-coding-test/internal/store"
)

type exoplanetService struct {
	store *store.StorageType
}

func NewExoplanetService() *exoplanetService {
	store := &store.StorageType{
		Mutex: sync.Mutex{},
	}

	return &exoplanetService{
		store: store,
	}
}

func (es *exoplanetService) AddExoplanet(exoplanet store.Exoplanet) {
	es.store.AddExoplanet(exoplanet)
}

func (es *exoplanetService) ListExoplanets() []store.Exoplanet {
	return es.store.ListExoplanets()
}

func (es *exoplanetService) GetExoplanetByID(id string) (store.Exoplanet, bool) {
	return es.store.GetExoplanetByID(id)
}

func (es *exoplanetService) UpdateExoplanet(exoplanet store.Exoplanet) bool {
	return es.store.UpdateExoplanet(exoplanet)
}

func (es *exoplanetService) DeleteExoplanet(id string) bool {
	return es.store.DeleteExoplanet(id)
}
