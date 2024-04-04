package pokemon

import (
	"sync"
)

type Pokedex struct {
	cache map[string]PokemonData
	mu    sync.Mutex
}

func NewPokedex() *Pokedex {
	px := Pokedex{
		cache: make(map[string]PokemonData),
	}

	return &px
}

func (pc *Pokedex) Add(key string, pd PokemonData) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cache[key] = pd
}

func (pc *Pokedex) Get(key string) (PokemonData, bool) {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	entry, found := pc.cache[key]
	if !found {
		return PokemonData{}, found
	}

	return entry, found
}
