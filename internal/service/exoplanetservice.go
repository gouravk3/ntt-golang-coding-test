package service

import (
	"sync"

	"github.com/gouravk3/ntt-golang-coding-test/internal/store"
)

type ExoplanetService struct {
	Store *store.StorageType
}

func NewExoplanetService() *ExoplanetService {
	store := &store.StorageType{
		Mutex: sync.Mutex{},
	}

	return &ExoplanetService{
		Store: store,
	}
}

func (es *ExoplanetService) AddExoplanet(exoplanet store.Exoplanet) {
	es.Store.AddExoplanet(exoplanet)
}

func (es *ExoplanetService) ListExoplanets() []store.Exoplanet {
	return es.Store.ListExoplanets()
}

func (es *ExoplanetService) GetExoplanetByID(id string) (store.Exoplanet, bool) {
	return es.Store.GetExoplanetByID(id)
}

func (es *ExoplanetService) UpdateExoplanet(exoplanet store.Exoplanet) bool {
	return es.Store.UpdateExoplanet(exoplanet)
}

func (es *ExoplanetService) DeleteExoplanet(id string) bool {
	return es.Store.DeleteExoplanet(id)
}
