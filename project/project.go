package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeAPIResponse struct {
	Id int `json:"id"`
	//Descriptions []string `json:"descriptions"`
}

var store = make(map[string]*PokeAPIResponse)

func main() {
	pokemon := "charizard"
	for range 5 {
		poke, err := GetPokemon(pokemon)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pokemon ID: %d\n", poke.Id)
	}
}

func GetPokemon(pokemon string) (*PokeAPIResponse, error) {
	if store[pokemon] != nil {
		return store[pokemon], nil
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	res, err := http.Get(url)
	fmt.Printf("Sent API request to %s\n", url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status not 200: %d", res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var pokedex PokeAPIResponse
	err = decoder.Decode(&pokedex)
	if err != nil {
		return nil, err
	}

	store[pokemon] = &pokedex
	return &pokedex, nil
}
