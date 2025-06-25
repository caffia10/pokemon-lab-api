package controller

type pokemonTypeDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type evolutionDto struct {
	Id          string `json:"id"`
	Requirement string `json:"requirement"`
}

type pokemonDto struct {
	Name       string           `json:"name"`
	Weight     string           `json:"weight"`
	Category   string           `json:"category"`
	ImgUrl     string           `json:"img_url"`
	Types      []pokemonTypeDto `json:"types"`
	Evolutions []evolutionDto
}
