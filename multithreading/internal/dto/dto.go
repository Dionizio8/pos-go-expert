package dto

import "github.com/Dionizio8/pos-go-expert/multithreading/internal/entity"

type BrasilAPIOutput struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (o *BrasilAPIOutput) ToEntity() entity.Address {
	return entity.Address{
		Cep:          o.Cep,
		State:        o.State,
		City:         o.City,
		Neighborhood: o.Neighborhood,
		Street:       o.Street,
	}
}

type ViacepAPIOutput struct {
	Cep         *string `json:"cep"`
	Logradouro  string  `json:"logradouro"`
	Complemento string  `json:"complemento"`
	Bairro      string  `json:"bairro"`
	Localidade  string  `json:"localidade"`
	Uf          string  `json:"uf"`
	Ibge        string  `json:"ibge"`
	Gia         string  `json:"gia"`
	Ddd         string  `json:"ddd"`
	Siafi       string  `json:"siafi"`
}

func (o *ViacepAPIOutput) ToEntity() entity.Address {
	return entity.Address{
		Cep:          *o.Cep,
		State:        o.Uf,
		City:         o.Localidade,
		Neighborhood: o.Bairro,
		Street:       o.Logradouro,
	}
}
