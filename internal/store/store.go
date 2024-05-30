package store

import (
	"sync"
)

type Exoplanet struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Distance    string `json:"distance"`
	Radius      string `json:"radius"`
	Mass        string `json:"mass,omitempty"`
	Type        string `json:"type"`
}

type StorageType struct {
	sync.Mutex
	exoplanets map[string]Exoplanet
	idCounter  int
}

func Init() *StorageType {
	return &StorageType{
		exoplanets: make(map[string]Exoplanet),
		idCounter:  1,
	}
}

func (s *StorageType) NextID() int {
	s.Lock()
	defer s.Unlock()
	id := s.idCounter
	s.idCounter++
	return id
}

func (s *StorageType) AddExoplanet(exoplanet Exoplanet) {
	s.Lock()
	defer s.Unlock()
	s.exoplanets[exoplanet.ID] = exoplanet
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
