package http

type pokemonDto struct {
	Name     string `json:"name"`
	Weight   string `json:"weight"`
	Category string `json:"category"`
	ImgUrl   string `json:"imgUrl"`
	Types    []*struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"types"`
}
