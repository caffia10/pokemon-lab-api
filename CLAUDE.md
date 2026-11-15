@~/dev/pokemon-lab-hub/ARCHITECTURE.md
@~/dev/pokemon-lab-hub/LEARNING_MODE.md

# pokedex (pokemon-lab-api)

## Responsabilidad

Fuente de verdad de **especies Pokémon y tipos**: alta, edición y consulta por id.

**No hace:** búsquedas complejas ni filtros de texto (eso es de `search`), ni nada de regiones,
gimnasios o profesores. Si un Pokémon necesita "saber" de qué región es, esa relación vive en
`geo`, no acá.

## Paradigma: Clean Architecture liviana en Go idiomático

- `internal/core` — entidades y **puertos** (interfaces de repositorio). Sin dependencias de
  frameworks ni de drivers.
- `internal/scylladb` — adaptadores que implementan los puertos.
- `internal/http` — handlers Fiber, DTOs y middlewares. Los handlers se construyen con
  funciones `makeXHandler(dep)` que cierran sobre sus dependencias.
- `internal/bootstrap` — wiring con `go.uber.org/dig`.
- `pkg/mderrors` — errores con metadata, reutilizables fuera del módulo.

Lo que se estudia acá, y por eso hay que respetarlo:
- Paquetes planos por capa, sin carpetas anidadas por dominio (se simplificó a propósito en
  los commits `b4142ed` y `e87f447`; no reintroducir anidamiento).
- Errores como valores: envolver con contexto y traducirlos a HTTP en
  `middlewares/errorHandler.go`, no en cada handler.
- Interfaces chicas y definidas donde se consumen. Lo idiomático es aceptar interfaces y
  devolver structs.
- Hay config de Mongo en `internal/config` que no se usa: Scylla es el storage real.

## Comandos

```bash
go run ./cmd/main.go          # requiere variables de envs/local.env (launch.json ya las carga)
go build ./... && go vet ./...
go test ./...
```

Requiere ScyllaDB en `localhost:9042` (ver `docker-compose.yml` en `poke-lab-ui`).

## Endpoints

`GET /ping` · `POST|PUT|GET /api/pokemons` · `GET /api/pokemons/:id` ·
`POST|GET /api/pokemon-types`
