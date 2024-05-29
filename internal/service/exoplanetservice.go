package service

type Exoplanet struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Distance    string `json:"distance"`
	Radius      string `json:"radius"`
	Mass        string `json:"mass,omitempty"`
	Type        string `json:"type"`
}

type exoplanetService struct {
}

func NewExoplanetService() *exoplanetService {
	return &exoplanetService{}
}
