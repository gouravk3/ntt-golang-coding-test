package store

import (
	"fmt"
	"sync"
)

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    int     `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass,omitempty"`
	Type        string  `json:"type"`
}

type StorageType struct {
	sync.Mutex
	exoplanets map[string]Exoplanet
	idCounter  int
}

func Init() *StorageType {
	return &StorageType{
		exoplanets: make(map[string]Exoplanet),
		idCounter:  0,
	}
}

func (s *StorageType) AddExoplanet(exoplanet Exoplanet) {
	s.Lock()
	defer s.Unlock()
	s.idCounter++
	id := fmt.Sprint(s.idCounter)
	exoplanet.ID = id
	s.exoplanets[id] = exoplanet
}

func (s *StorageType) ListExoplanets() []Exoplanet {
	s.Lock()
	defer s.Unlock()
	planets := make([]Exoplanet, 0, len(s.exoplanets))
	for _, exoplanet := range s.exoplanets {
		planets = append(planets, exoplanet)
	}
	return planets
}

func (s *StorageType) GetExoplanetByID(id string) (Exoplanet, bool) {
	s.Lock()
	defer s.Unlock()
	exoplanet, found := s.exoplanets[id]
	return exoplanet, found
}

func (s *StorageType) UpdateExoplanet(exoplanet Exoplanet) bool {
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
