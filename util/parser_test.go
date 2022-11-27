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
