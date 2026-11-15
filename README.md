# pokedex — pokemon-lab-api

Fuente de verdad de especies Pokémon y tipos.

**Stack:** Go · Fiber v2 · ScyllaDB · dig — **Paradigma:** Clean Architecture liviana.

```bash
go run ./cmd/main.go   # con envs/local.env
```

Arquitectura, convenciones y deuda: [CLAUDE.md](CLAUDE.md) · Mapa del sistema: `~/dev/pokemon-lab-hub/ARCHITECTURE.md`

## INSTALLATION REQUIREMENTS:

- install cqlsh
- setup scylladb node in some docker and map his port
- map .envs

## Tips

### Check data

- open terminal
- cqlsh localhost 9042
- USE {key_space}; e.g -> USE pokemon_space;
- SELECT \* FROM {table}; -> e.g: SELECT \* FROM pokemons;
- expected result will showed like:

| id                                   | category | img_url                                                             | name    | types        | weight |
| ------------------------------------ | -------- | ------------------------------------------------------------------- | ------- | ------------ | ------ |
| 37245c80-ffe7-4b3c-828a-59bd1b268273 | Mouse    | https://img.pokemondb.net/sprites/scarlet-violet/normal/pikachu.png | Pikachu | ['Electric'] | 2.5lbs |
