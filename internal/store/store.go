package store

import (
    "sync"

	"github.com/gouravk3/ntt-golang-coding-test/internal/service"
)

type StorageType struct {
    sync.Mutex
    exoplanets map[string]service.Exoplanet
    idCounter  int
}

var Storage = StorageType{
    exoplanets: make(map[string]service.Exoplanet),
    idCounter:  1,
}

func (s *StorageType) NextID() int {
    s.Lock()
    defer s.Unlock()
    id := s.idCounter
    s.idCounter++
    return id
}

func (s *StorageType) AddExoplanet(exoplanet service.Exoplanet) {
    s.Lock()
    defer s.Unlock()
    s.exoplanets[exoplanet.ID] = exoplanet
}

func (s *StorageType) ListExoplanets() []service.Exoplanet {
    s.Lock()
    defer s.Unlock()
    planets := make([]service.Exoplanet, 0, len(s.exoplanets))
    for _, exoplanet := range s.exoplanets {
        planets = append(planets, exoplanet)
    }
    return planets
}

func (s *StorageType) GetExoplanetByID(id string) (service.Exoplanet, bool) {
    s.Lock()
    defer s.Unlock()
    exoplanet, found := s.exoplanets[id]
    return exoplanet, found
}

func (s *StorageType) UpdateExoplanet(exoplanet service.Exoplanet) bool {
    s.Lock()
    defer s.Unlock()
    if _, found := s.exoplanets[exoplanet.ID]; found {
        s.exoplanets[exoplanet.ID] = exoplanet
        return true
    }
    return false
}

func (s *StorageType) DeleteExoplanet(id string) bool {
    s.Lock()
    defer s.Unlock()
    if _, found := s.exoplanets[id]; found {
        delete(s.exoplanets, id)
        return true
    }
    return false
}
