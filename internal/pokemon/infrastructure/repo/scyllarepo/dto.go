package scyllarepo

type pokemonDto struct {
	Id       string   `db:"id"`
	Name     string   `db:"name"`
	Category string   `db:"category"`
	Weight   string   `db:"weight"`
	ImgUrl   string   `db:"img_url"`
	Types    []string `db:"types"` // save only the id
}
