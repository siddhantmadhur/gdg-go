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

func main() {
	pokemon := "charizard"
	_, err := GetPokemon(pokemon)
	if err != nil {
		panic(err)
	}
}

func GetPokemon(pokemon string) (string, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("Status not 200: %d", res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var pokedex PokeAPIResponse
	err = decoder.Decode(&pokedex)
	if err != nil {
		return "", err
	}

	fmt.Printf("ID: %d\n", pokedex.Id)

	return "", nil
}
