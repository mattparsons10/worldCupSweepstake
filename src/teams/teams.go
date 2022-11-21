package teams

import (
	"fmt"
	"math/rand"
	"time"
)

func AllocateTeamsToPlayers(players []string, availableTeams []string) map[string][]string {

	playerSelections := make(map[string][]string)

	lastTeam := false

	for len(availableTeams) > 0 && !lastTeam {
		for _, player := range players {
			if lastTeam {
				break
			}
			fmt.Println("Number of teams left in the draw ", len(availableTeams))

			min := 0
			max := len(availableTeams) - 1

			randomTeamNumber := 0

			if max > 0 {
				randomTeamNumber = generateRandomNumber(min, max)
			}
			fmt.Printf(fmt.Sprintf("Players: %s has drawn the team %s", player, availableTeams[randomTeamNumber]))

			playerSelections[player] = append(playerSelections[player], availableTeams[randomTeamNumber])

			if len(availableTeams) == 1 {
				lastTeam = true
			} else {
				availableTeams = stringSliceDelete(availableTeams, randomTeamNumber) // Truncate slice.
			}
		}
	}
	return playerSelections
}

func FetchWorldCupTeams() []string {

	return []string{"Qatar", "Germany", "Denmark", "France", "Belgium", "Croatia", "Spain", "Serbia", "England", "Switzerland", "Netherlands", "Brazil", "Argentina", "Ecuador", "Uruguay", "Iran", "South Korea", "Japan", "Saudi Arabia", "Canada", "Poland", "Portugal", "Senegal", "Ghana", "Morocco", "Tunisia", "Cameroon", "United States", "Mexico", "Wales", "Australia", "Costa Rica"}

}

func stringSliceDelete(slice []string, i int) []string {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
