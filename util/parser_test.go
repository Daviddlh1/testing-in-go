package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

// If you want to run only one of the tests  in the testing file you need to use the command:
// go test ./util -run=Test<name of the test function here> -v
// Example: go test ./util -run=TestParserPokemonSuccess -v
func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	parsedPokemon, err := ParsePokemon(response)
	c.NoError(err)

	body, err = ioutil.ReadFile("samples/api_response.json")
	c.NoError(err)

	var expectedPokemon models.Pokemon
	err = json.Unmarshal([]byte(body), &expectedPokemon)

	c.Equal(expectedPokemon, parsedPokemon)
}

func TestParserPokemonTypeNotFound(t *testing.T) {
	c := require.New(t)
	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	response.PokemonType = []models.PokemonType{}

	_, err = ParsePokemon(response)
	c.NotNil(err)
	c.EqualError(ErrNotFoundPokemonType, err.Error())
}

func TestParserPokemonTypeNameNotFound(t *testing.T) {
	c := require.New(t)
	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)
	var pokemonType = models.PokemonType{}
	response.PokemonType = []models.PokemonType{pokemonType}

	_, err = ParsePokemon(response)
	c.NotNil(err)
	c.EqualError(ErrNotFoundPokemonTypeName, err.Error())
}

// Command to run benchmarks: go test ./util -run=Parser -bench=bench.old
func BenchmarkParser(b *testing.B) {
	c := require.New(b)

	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	for n := 0; n < b.N; n++ {
		_, err := ParsePokemon(response)
		c.NoError(err)
	}

}
