package teams

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
	"golang.org/x/exp/slices"
)

func TestFetchWorldCupTeamsLength(t *testing.T) {
	teams := FetchWorldCupTeams()

	assert.Equal(t, len(teams), 32)
}

func TestAllocateTeamsToPlayers(t *testing.T) {
	testCases := []struct {
		desc                     string
		listOfPlayers            []string
		expectedResponse         map[string][]string
		expectedNoTeamsPerPlayer int
	}{
		{
			desc:                     "1 player gets all of the teams assigned to them",
			listOfPlayers:            []string{"Player1"},
			expectedResponse:         map[string][]string{},
			expectedNoTeamsPerPlayer: 32,
		},
		{
			desc:                     "32 players get one of the teams assigned to them",
			listOfPlayers:            []string{"Player1", "Player2", "Player3", "Player4", "Player5", "Player6", "Player7", "Player8", "Player9", "Player10", "Player11", "Player12", "Player13", "Player14", "Player15", "Player16", "Player17", "Player18", "Player19", "Player20", "Player21", "Player22", "Player23", "Player24", "Player25", "Player26", "Player27", "Player28", "Player29", "Player30", "Player31", "Player32"},
			expectedResponse:         map[string][]string{},
			expectedNoTeamsPerPlayer: 1,
		},
		{
			desc:                     "16 players get 2 of the teams assigned to them",
			listOfPlayers:            []string{"Player1", "Player2", "Player3", "Player4", "Player5", "Player6", "Player7", "Player8", "Player9", "Player10", "Player11", "Player12", "Player13", "Player14", "Player15", "Player16"},
			expectedResponse:         map[string][]string{},
			expectedNoTeamsPerPlayer: 2,
		},
		{
			desc:                     "4 players get 8 of the teams assigned to them",
			listOfPlayers:            []string{"Player1", "Player2", "Player3", "Player4"},
			expectedResponse:         map[string][]string{},
			expectedNoTeamsPerPlayer: 8,
		},
	}

	for _, tC := range testCases {
		res := AllocateTeamsToPlayers(tC.listOfPlayers, FetchWorldCupTeams())

		assert.Equal(t, len(res), len(tC.listOfPlayers))
		assert.Equal(t, len(res["Player1"]), tC.expectedNoTeamsPerPlayer)
	}
}

func TestFetchWorldCupTeamsContainsCountry(t *testing.T) {

	testCases := []struct {
		desc       string
		country    string
		isReturned bool
	}{
		{
			desc:       "England are one of the teams in the world cup",
			country:    "England",
			isReturned: true,
		},
		{
			desc:       "Scotland are not one of the teams in the world cup",
			country:    "Scotland",
			isReturned: false,
		},
		{
			desc:       "Brazil are one of the teams in the world cup",
			country:    "Brazil",
			isReturned: true,
		},
		{
			desc:       "Poland are one of the teams in the world cup",
			country:    "Poland",
			isReturned: true,
		},
		{
			desc:       "Ghana are one of the teams in the world cup",
			country:    "Ghana",
			isReturned: true,
		},
		{
			desc:       "China are not one of the teams in the world cup",
			country:    "China",
			isReturned: false,
		},
	}

	for _, tC := range testCases {

		teams := FetchWorldCupTeams()

		fmt.Println(teams, tC.country)

		assert.Equal(t, slices.Contains(teams, tC.country), tC.isReturned)
	}

}
