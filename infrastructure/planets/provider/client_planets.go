package provider

import "context"

type GetPlanetsResponse struct {
	Name    string
	Climate string
	Terrain string
}

func (e *SwApi) GetPlanets(ctx context.Context) *GetPlanetsResponse {
}
